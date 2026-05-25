# calculadora.py

def sumar(a, b):
    return a + b

def restar(a, b):
    return a - b

def multiplicar(a, b):
    return a * b

def dividir(a, b):
    if b == 0:
        return "Error: no se puede dividir entre 0"
    return a / b


while True:
    print("\n===== CALCULADORA =====")
    print("1. Sumar")
    print("2. Restar")
    print("3. Multiplicar")
    print("4. Dividir")
    print("5. Salir")

    opcion = input("Selecciona una opción: ")

    if opcion == "5":
        print("Saliendo de la calculadora...")
        break

    if opcion not in ["1", "2", "3", "4"]:
        print("Opción inválida")
        continue

    try:
        num1 = float(input("Ingresa el primer número: "))
        num2 = float(input("Ingresa el segundo número: "))

        if opcion == "1":
            resultado = sumar(num1, num2)
            print(f"Resultado: {resultado}")

        elif opcion == "2":
            resultado = restar(num1, num2)
            print(f"Resultado: {resultado}")

        elif opcion == "3":
            resultado = multiplicar(num1, num2)
            print(f"Resultado: {resultado}")

        elif opcion == "4":
            resultado = dividir(num1, num2)
            print(f"Resultado: {resultado}")

    except ValueError:
        print("Error: debes ingresar números válidos")