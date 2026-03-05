# Go-Fisio-Concurrencia 🏊‍♂️ | Gestión de Recursos con Concurrencia en Go

[![Go Report Card](https://goreportcard.com/badge/github.com/Gabo-Dev/go-fisio-concurrencia)](https://goreportcard.com/report/github.com/Gabo-Dev/go-fisio-concurrencia)
[![Go Reference](https://pkg.go.dev/badge/github.com/Gabo-Dev/go-fisio-concurrencia.svg)](https://pkg.go.dev/badge/github.com/Gabo-Dev/go-fisio-concurrencia)

## Introducción

**Go-Fisio-Concurrencia** es una simulación de escritorio desarrollada en Go que demuestra el uso de patrones de concurrencia avanzados como goroutines, canales y `sync.WaitGroup`. Su objetivo principal es simular la gestión de recursos (fisioterapeutas, masajistas) y la distribución de clientes con prioridad en un centro de fisioterapia.

Este proyecto nació de mi interés por aprender y dominar un lenguaje moderno y sólido como **Go**, un "safe place" no tan alterado por la influencia constante de nuevas tecnologías. Este impulso se vio reforzado al ver una clase de Gentleman y notar las similitudes de Go con los lenguajes utilizados en mis prácticas universitarias de programación concurrente y distribuida. Mi objetivo fue profundizar en los conceptos de Go, traduciendo una de mis prácticas a este nuevo dominio para entender mejor su implementación.

**Nota del Desarrollador:** Este proyecto fue creado meramente con el propósito de aprender y explorar a fondo los conceptos de concurrencia en Go. Aunque funcional, su diseño primario no fue pensando en la contribución externa, sino como una herramienta personal de estudio y experimentación.

## ✨ Características Clave

*   **Simulación de Centro de Fisioterapia:** Modela la operación de un centro con diferentes tipos de clientes (rehabilitación, masajes) y recursos (fisioterapeutas, masajistas).
*   **Gestión de Concurrencia Avanzada:** Utiliza goroutines y canales para manejar múltiples clientes y empleados simultáneamente.
*   **Despachador con Prioridad:** Implementa un sistema de despacho que asigna recursos dando prioridad a clientes de rehabilitación.
*   **Manejo de Recursos con Semáforos:** Controla la disponibilidad de fisioterapeutas y masajistas utilizando canales buffer como semáforos.
*   **Flujo de Clientes Definido:** Cada cliente sigue un ciclo de vida desde su llegada, espera, atención y salida, liberando recursos.
*   **Patrones de Concurrencia Idiomáticos de Go:** Demuestra la aplicación práctica de `sync.WaitGroup`, `select`, y la inversión de dependencias mediante interfaces.

## 🛠️ Tecnologías Utilizadas

*   **Lenguaje**: Go
*   **Concurrencia**: Goroutines, Channels, `sync.WaitGroup`
*   **Estructura**: `cmd/` e `internal/` para modularidad
*   **Herramientas**: `go.mod` para gestión de módulos

## 🚀 Primeros Pasos

Sigue estos pasos para configurar y ejecutar la simulación de Go-Fisio-Concurrencia en tu entorno local:

### Requisitos Previos

Asegúrate de tener instalado Go en tu sistema. Puedes descargarlo desde [golang.org](https://golang.org/dl/).

### Instalación y Ejecución

1.  **Clona el repositorio:**
    ```bash
    git clone https://github.com/Gabo-Dev/go-fisio-concurrencia.git
    cd go-fisio-concurrencia
    ```

2.  **Ejecuta la simulación:**
    ```bash
    go run cmd/fisio-center/main.go
    ```
    Esto iniciará la simulación y verás el flujo de clientes y la gestión de recursos en tu terminal.

## 💡 Reflexiones del Desarrollador y Viaje de Aprendizaje

Este proyecto ha sido una inmersión profunda en el mundo de la concurrencia en Go y ha representado una curva de aprendizaje significativa.

### Desafíos Superados

El mayor reto fue el **primer cara a cara con las goroutines y los tecnicismos de la concurrencia en Go**. Aunque la lógica conceptual detrás de los sistemas concurrentes me resultaba familiar por mis estudios universitarios, asimilar la magnitud y las implicaciones prácticas de Go (cómo se comportan las goroutines, la seguridad de los canales, la coordinación con `sync.WaitGroup`) fue un desafío constante. Requirió experimentar, depurar y leer mucho para entender no solo cómo funcionan, sino *por qué* funcionan de esa manera y cómo aplicarlos correctamente para evitar condiciones de carrera o bloqueos.

### Aprendizajes Clave y Crecimiento

Este proyecto me brindó una tremenda experiencia de aprendizaje, consolidando mi crecimiento como desarrollador Go. Algunos de los puntos clave incluyen:

*   **Dominio de la Concurrencia en Go:** Comprensión profunda de goroutines, canales y `sync.WaitGroup` en escenarios prácticos.
*   **Refuerzo de Conceptos Fundamentales:** Sirvió para reforzar conceptos de programación de bajo nivel como hilos (análogos a goroutines en este contexto) y el manejo de punteros, que son cruciales en Go.
*   **Aplicación Práctica de Patrones:** Entender cómo implementar patrones de concurrencia como semáforos, productor-consumidor y despachadores con prioridad utilizando las herramientas idiomáticas de Go.
*   **Diseño de Software:** Aprender a estructurar un proyecto Go con paquetes `cmd/` e `internal/` para una mejor modularidad y cómo usar interfaces para la Inversión de Dependencias, rompiendo ciclos de importación.

## 🤝 Contribuciones

Este proyecto fue desarrollado principalmente como una herramienta personal de estudio y experimentación para aprender y explorar a fondo los conceptos de concurrencia en Go. Por esta razón, su diseño primario no fue pensando en la contribución externa. Sin embargo, si encuentras un error o tienes alguna sugerencia de mejora, no dudes en abrir un "issue".

## 📧 Contacto

Si tienes alguna pregunta, sugerencia o simplemente quieres conectar, ¡no dudes en contactarme!

*   **GitHub**: [@Gabo-Dev](https://github.com/Gabo-Dev)