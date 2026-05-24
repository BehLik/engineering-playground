package main
 
import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"
)
 
// Resultado agrupa el valor y el operador usado
type Resultado struct {
    Valor float64
    Operador string
}
 
// calcular recibe dos números y un operador, devuelve (resultado, error)
func calcular(a, b float64, op string) (Resultado, error) {
    switch op {
    case "+":
        return Resultado{a + b, op}, nil
    case "-":
        return Resultado{a - b, op}, nil
    case "*":
        return Resultado{a * b, op}, nil
    case "/":
        if b == 0 {
            return Resultado{}, errors.New("división entre cero")
        }
        return Resultado{a / b, op}, nil
    default:
        return Resultado{}, errors.New("operador desconocido")
    }
}
 
// leerLinea lee una línea de texto del usuario
func leerLinea(reader *bufio.Reader) (string, error) {
    texto, err := reader.ReadString('\\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(texto), nil
}
 
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("=== Calculadora en Go ===")
 
    for {
        fmt.Print("Ingresa: número operador número (ej: 10 / 2): ")
 
        linea, err := leerLinea(reader)
        if err != nil { break }
        if linea == "salir" { break }
 
        partes := strings.Fields(linea)
        if len(partes) != 3 {
            fmt.Println("Formato incorrecto")
            continue
        }
 
        a, errA := strconv.ParseFloat(partes[0], 64)
        b, errB := strconv.ParseFloat(partes[2], 64)
        if errA != nil || errB != nil {
            fmt.Println("Números inválidos")
            continue
        }
 
        res, err := calcular(a, b, partes[1])
        if err != nil {
            fmt.Printf("Error: %v\\n", err)
            continue
        }
        fmt.Printf("Resultado (%s): %.2f\\n", res.Operador, res.Valor)
    }
    fmt.Println("¡Hasta luego!")
}