package strategies

const (
    StrategyPercentage = "percentage"
    StrategySomethingelse = "lol"
)

func GetStrategyTypes() ([]string, error) {
    strategies := []string{StrategyPercentage, StrategySomethingelse}
    return strategies, nil
}
