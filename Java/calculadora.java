package Java;
import java.util.Scanner;
public class calculadora {

    public static double sumar(double a, double b) {
        return a + b;
    }

    public static double restar(double a, double b) {
        return a - b;
    }

    public static double multiplicar(double a, double b) {
        return a * b;
    }

    public static double dividir(double a, double b) {
        if (b == 0) {
            System.out.println("Error: no se puede dividir entre 0");
            return 0;
        }
        return a / b;
    }

    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);

        while (true) {

            System.out.println("\n===== CALCULADORA =====");
            System.out.println("1. Sumar");
            System.out.println("2. Restar");
            System.out.println("3. Multiplicar");
            System.out.println("4. Dividir");
            System.out.println("5. Salir");

            System.out.print("Selecciona una opción: ");
            int opcion = scanner.nextInt();

            if (opcion == 5) {
                System.out.println("Saliendo de la calculadora...");
                break;
            }

            if (opcion < 1 || opcion > 5) {
                System.out.println("Opción inválida");
                continue;
            }

            System.out.print("Ingresa el primer número: ");
            double num1 = scanner.nextDouble();

            System.out.print("Ingresa el segundo número: ");
            double num2 = scanner.nextDouble();

            double resultado = 0;

            switch (opcion) {

                case 1:
                    resultado = sumar(num1, num2);
                    break;

                case 2:
                    resultado = restar(num1, num2);
                    break;

                case 3:
                    resultado = multiplicar(num1, num2);
                    break;

                case 4:
                    resultado = dividir(num1, num2);
                    break;
            }

            System.out.println("Resultado: " + resultado);
        }

        scanner.close();
    }
}