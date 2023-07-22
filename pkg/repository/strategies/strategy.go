package strategies

const (
    StrategyPercentage = "percentage"
    StrategySomethingelse = "somethingelsestrategy"
)

func GetStrategyTypes() ([]string, error) {
    strategies := []string{StrategyPercentage, StrategySomethingelse}
    return strategies, nil
}
