import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Logger logger = new Logger();
        Calculator calculator = new Calculator(logger);

        System.out.print("Введите действительную часть первого комплексного числа: ");
        double real1 = scanner.nextDouble();
        System.out.print("Введите мнимую часть первого комплексного числа: ");
        double imaginary1 = scanner.nextDouble();

        System.out.print("Введите действительную часть второго комплексного числа: ");
        double real2 = scanner.nextDouble();
        System.out.print("Введите мнимую часть второго комплексного числа: ");
        double imaginary2 = scanner.nextDouble();

        ComplexNumber a = new ComplexNumber(real1, imaginary1);
        ComplexNumber b = new ComplexNumber(real2, imaginary2);

        ComplexNumber sum = calculator.add(a, b);
        System.out.println("Сумма: " + sum);
        ComplexNumber subtract = calculator.subtract(a, b);
        System.out.println("Вычитание: " + subtract);
        ComplexNumber product = calculator.multiply(a, b);
        System.out.println("Произведение: " + product);
        ComplexNumber quotient = calculator.divide(a, b);
        System.out.println("Частное: " + quotient);


        scanner.close();
    }
}