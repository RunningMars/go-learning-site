package services

import (
	"errors"
	"time"

	"github.com/RunningMars/go-learning-site/backend/models"
	"github.com/RunningMars/go-learning-site/backend/utils"
)

type UserService struct {
	repo  *models.UserRepository
	redis *utils.RedisClient
}

func NewUserService(repo *models.UserRepository, redis *utils.RedisClient) *UserService {
	return &UserService{
		repo:  repo,
		redis: redis,
	}
}

func (s *UserService) Register(req *models.RegisterRequest) (*models.User, error) {
	// 检查用户名是否已存在
	exists, err := s.repo.CheckUsernameExists(req.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 检查手机号是否已存在
	exists, err = s.repo.CheckPhoneExists(req.Phone)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("手机号已被注册")
	}

	// 创建用户
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Phone:    req.Phone,
		Email:    req.Email,
		Age:      req.Age,
		Gender:   req.Gender,
	}

	return s.repo.CreateUser(user, passwordHash)
}

func (s *UserService) Login(phone, password string) (string, error) {
	user, err := s.repo.GetUserByPhone(phone)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("用户不存在")
	}

	// 验证密码
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("密码错误")
	}

	// 生成 JWT token
	return utils.GenerateJWTToken(user.ID)
}

func (s *UserService) SendVerificationCode(phone string) error {
	// 生成6位随机验证码
	code := utils.GenerateRandomCode(6)

	// 将验证码保存到 Redis，设置5分钟过期
	key := "verification_code:" + phone
	err := s.redis.Set(key, code, 5*time.Minute)
	if err != nil {
		return err
	}

	// TODO: 调用短信服务发送验证码
	return nil
}

func (s *UserService) VerifyCode(phone, code string) bool {
	key := "verification_code:" + phone
	storedCode, err := s.redis.Get(key)
	if err != nil {
		return false
	}

	return storedCode == code
}
