public class Calculator {
    private Logger logger;

    public Calculator(Logger logger) {
        this.logger = logger;
    }

    public ComplexNumber add(ComplexNumber a, ComplexNumber b) {
        logger.log("Выполнение сложения: " + a + " + " + b);
        return a.add(b);
    }

    public ComplexNumber subtract(ComplexNumber a, ComplexNumber b) {
        logger.log("Выполнение вычитания: " + a + " - " + b);
        return a.subtract(b);
    }

    public ComplexNumber multiply(ComplexNumber a, ComplexNumber b) {
        logger.log("Выполнение умножения: " + a + " * " + b);
        return a.multiply(b);
    }

    public ComplexNumber divide(ComplexNumber a, ComplexNumber b) {
        logger.log("Выполнение деления: " + a + " / " + b);
        return a.divide(b);
    }
}