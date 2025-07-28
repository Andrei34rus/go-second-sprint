package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат данных: требуется 3 параметра")
	}

	// Обработка шагов
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return errors.New("неверный формат количества шагов")
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}

	// Обработка типа тренировки
	trainingType := strings.TrimSpace(parts[1])
	if trainingType == "" {
		return errors.New("тип тренировки не может быть пустым")
	}

	// Обработка длительности
	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return errors.New("неверный формат длительности")
	}
	if duration <= 0 {
		return errors.New("длительность должна быть положительной")
	}

	// Сохраняем значения в структуру
	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = duration

	return nil
}
func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var calories float64
	var err error
	switch strings.ToLower(t.TrainingType) {
	case "бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки %v", err)
	}
	if err != nil {
		return "", fmt.Errorf("ошибка расчёта калорий: %v", err)
	}
	hours := t.Duration.Hours()
	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, hours, distance, speed, calories)

	return info, nil
}
