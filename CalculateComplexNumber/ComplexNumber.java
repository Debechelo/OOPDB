public class ComplexNumber {
    private double real;
    private double imaginary;

    public ComplexNumber(double real, double imaginary) {
        this.real = real;
        this.imaginary = imaginary;
    }

    public ComplexNumber add(ComplexNumber other) {
        return new ComplexNumber(this.real + other.real, this.imaginary + other.imaginary);
    }

    public ComplexNumber subtract(ComplexNumber other) {
        return new ComplexNumber(this.real - other.real, this.imaginary - other.imaginary);
    }

    public ComplexNumber multiply(ComplexNumber other) {
        double real = this.real * other.real - this.imaginary * other.imaginary;
        double imaginary = this.real * other.imaginary + this.imaginary * other.real;
        return new ComplexNumber(real, imaginary);
    }

    public ComplexNumber divide(ComplexNumber other) {
        double denominator = other.real * other.real + other.imaginary * other.imaginary;
        double real = (this.real * other.real + this.imaginary * other.imaginary) / denominator;
        double imaginary = (this.imaginary * other.real - this.real * other.imaginary) / denominator;
        return new ComplexNumber(real, imaginary);
    }

    @Override
    public String toString() {
        if(imaginary == 0.0)
            return String.format("%.1f", real);
        return String.format("%.1f + %.1fi", real, imaginary);
    }
}