# Go-Fisio-Concurrencia 🏊‍♂️

**Go-Fisio-Concurrencia** es una simulación de escritorio para explorar y dominar patrones de concurrencia avanzados en Go, sirviendo como un caso de estudio práctico en ingeniería y arquitectura de software.

---

## 🧠 Filosofía del Proyecto

Este repositorio es más que código; es un manifiesto sobre el aprendizaje profundo y la maestría de herramientas.

- **Origen:** La extrapolación de un problema de concurrencia universitario (originalmente en Java) a una solución idiomática en Go, enfocándose en el "porqué" de los patrones del lenguaje.
- **Objetivo Principal:** Utilizar el proyecto como un vehículo para dominar conceptos avanzados de **lenguaje (Go)**, **ingeniería de software (buenas prácticas, testing)** y **arquitectura de software (diseño limpio)**.
- **Metodología:** Un compromiso con el uso de herramientas profesionales como **Neovim** y **Fish Shell** para forjar eficiencia y un entendimiento fundamental del entorno de desarrollo.
- **Proceso de Aprendizaje:** Seguir un modelo de mentoría con **OpenCode**, donde un agente de IA actúa como Arquitecto Senior para guiar el diseño y el razonamiento, promoviendo el desarrollo de un criterio de ingeniería en lugar de la simple generación de código.

---

## 🎯 El Desafío Técnico

La simulación modela un centro de fisioterapia con varios retos de concurrencia:

- **Gestión de Recursos**: Implementación de semáforos y mutex usando canales para controlar el acceso a empleados y vestuarios.
- **Lógica de Prioridad**: Diseño de un sistema donde ciertos clientes (Rehabilitación) tienen prioridad de acceso.
- **Sincronización Compleja**: Asegurar que los empleados queden bloqueados hasta que sus clientes completen todo el ciclo (sesión + vestuario).

---

## 🛠️ Stack de Desarrollo

| Herramienta    | Propósito                                                                                            |
| :------------- | :--------------------------------------------------------------------------------------------------- |
| **Go**         | Lenguaje núcleo del proyecto                                                                         |
| **Neovim**     | Editor configurado con [Gentleman Dots](https://github.com/Gentleman-Programming/Gentleman-dotfiles) |
| **OpenCode**   | Mentoría y arquitectura guiada por IA                                                                |
| **Fish Shell** | Entorno de terminal y productividad                                                                  |
| **Git**        | Control de versiones                                                                                 |

---

## 🏗️ Arquitectura y Conceptos de Go

### Conceptos Clave Aplicados

- **Goroutines & Channels**: Como pilares para la concurrencia, comunicación y sincronización.
- **`sync.WaitGroup`**: Para orquestar la finalización ordenada de la simulación.
- **Estructura de Paquetes**: Layout `cmd/` e `internal/` para una clara separación de responsabilidades.

### Estructura de Directorios

```text
src/
├── cmd/
│   └── fisio-center/   # Proceso principal de la simulación (main.go)
└── internal/
    ├── centro/         # Lógica de negocio y sincronización del centro
    └── cliente/        # Definición y ciclo de vida de los clientes
```

💡 Conceptos Clave de Go Aplicados
Este proyecto sirve como referencia práctica para los siguientes patrones y conceptos de Go:

- **`sync.WaitGroup`**:
  - **Uso**: Para sincronizar la finalización de un grupo de goroutines.
  - **Patrón**: `Wg.Add(1)` se llama antes de lanzar cada goroutine. `defer Wg.Done()` se coloca al inicio de la función de la goroutine para asegurar que se reste del contador al terminar. `Wg.Wait()` en la goroutine principal se bloquea hasta que el contador llega a cero.
- **Canales como Semáforos (`chan struct{}, N`)**:
  - **Uso**: Limitar el acceso concurrente a un pool de `N` recursos.
  - **Patrón**: Se crea con `make(chan struct{}, N)`. Una goroutine "toma" un recurso con `<-canal` (se bloquea si no hay) y lo "devuelve" con `canal <- struct{}{}`.
- **`select` para Competencia de Recursos**:
  - **Uso**: Permitir a una goroutine esperar por varios canales a la vez y actuar sobre el primero que esté disponible.
  - **Patrón**: El cliente de masaje lo usa para esperar por `MasajistaDisponible` **O** `FisioDisponible`, tomando el que se libere primero.
- **Patrón Dispatcher con Prioridad**:
  - **Uso**: Centralizar la lógica de asignación de recursos cuando hay reglas de negocio complejas (ej. prioridad).
  - **Patrón**: Una goroutine (`runDispatcher`) en un bucle `for-select` que, gracias a `select` anidados con `default`, siempre chequea la cola de alta prioridad (`ColaRehabilitacion`) antes que las demás.
- **Interfaces para Inversión de Dependencias**:
  - **Uso**: Para desacoplar paquetes y romper ciclos de importación (el error `import cycle not allowed`).
  - **Patrón**: El paquete de nivel inferior (`cliente`) define una `interface` con los métodos que necesita (`CenterService`). El paquete de nivel superior (`centro`) implementa esa interfaz, cumpliendo el contrato sin crear una dependencia circular.
- **Rutas de Importación de Módulos**:
  - **Uso**: Entender por qué los `import` en Go son absolutos (ej. `github.com/user/repo/...`).
  - **Patrón**: El nombre base de los imports se define en `go.mod`. Go no usa imports relativos como `../` para evitar ambigüedades y hacer el código más robusto.
