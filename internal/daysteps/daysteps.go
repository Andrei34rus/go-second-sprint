package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	stepLengthCoefficient = 0.45 // Коэффициент длины шага
	mInKm                 = 1000 // Метров в километре
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат данных")
	}
	stepsStr := parts[0]
	if strings.HasPrefix(stepsStr, " ") || strings.HasSuffix(stepsStr, " ") {
		return errors.New("неверный формат количества шагов")
	}
	if strings.Contains(stepsStr, " ") {
		return errors.New("неверный формат количества шагов")
	}
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return errors.New("неверный формат количества шагов")
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}
	durationStr := parts[1]
	if strings.HasPrefix(durationStr, " ") || strings.HasSuffix(durationStr, " ") {
		return errors.New("неверный формат длительности")
	}

	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return errors.New("неверный формат длительности")
	}
	if duration <= 0 {
		return errors.New("длительность должна быть больше 0")
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}
func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if ds.Steps <= 0 {
		return "", errors.New("количество шагов должно быть больше 0")
	}
	if ds.Duration <= 0 {
		return "", errors.New("длительность должна быть больше 0")
	}
	if ds.Weight <= 0 {
		return "", errors.New("вес должен быть больше 0")
	}
	if ds.Height <= 0 {
		return "", errors.New("рост должен быть больше 0")
	}
	stepLength := ds.Height * stepLengthCoefficient
	distance := float64(ds.Steps) * stepLength / mInKm
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", fmt.Errorf("ошибка расчёта калорий: %w", err)
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
