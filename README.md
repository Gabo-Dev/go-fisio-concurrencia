# Go-Fisio-Concurrencia 🏊‍♂️

**Go-Fisio-Concurrencia** es una simulación de escritorio para explorar y dominar patrones de concurrencia avanzados en Go, sirviendo como un caso de estudio práctico en ingeniería y arquitectura de software.

---

## 🧠 Filosofía del Proyecto

Este repositorio es más que código; es un manifiesto sobre el aprendizaje profundo y la maestría de herramientas.

* **Origen:** La extrapolación de un problema de concurrencia universitario (originalmente en Java) a una solución idiomática en Go, enfocándose en el "porqué" de los patrones del lenguaje.
* **Objetivo Principal:** Utilizar el proyecto como un vehículo para dominar conceptos avanzados de **lenguaje (Go)**, **ingeniería de software (buenas prácticas, testing)** y **arquitectura de software (diseño limpio)**.
* **Metodología:** Un compromiso con el uso de herramientas profesionales como **Neovim** y **Fish Shell** para forjar eficiencia y un entendimiento fundamental del entorno de desarrollo.
* **Proceso de Aprendizaje:** Seguir un modelo de mentoría con **OpenCode**, donde un agente de IA actúa como Arquitecto Senior para guiar el diseño y el razonamiento, promoviendo el desarrollo de un criterio de ingeniería en lugar de la simple generación de código.

---

## 🎯 El Desafío Técnico

La simulación modela un centro de fisioterapia con varios retos de concurrencia:

* **Gestión de Recursos**: Implementación de semáforos y mutex usando canales para controlar el acceso a empleados y vestuarios.
* **Lógica de Prioridad**: Diseño de un sistema donde ciertos clientes (Rehabilitación) tienen prioridad de acceso.
* **Sincronización Compleja**: Asegurar que los empleados queden bloqueados hasta que sus clientes completen todo el ciclo (sesión + vestuario).

---

## 🛠️ Stack de Desarrollo

| Herramienta | Propósito |
| :--- | :--- |
| **Go** | Lenguaje núcleo del proyecto |
| **Neovim** | Editor configurado con [Gentleman Dots](https://github.com/Gentleman-Programming/Gentleman-dotfiles) |
| **OpenCode** | Mentoría y arquitectura guiada por IA |
| **Fish Shell** | Entorno de terminal y productividad |
| **Git** | Control de versiones |

---

## 🏗️ Arquitectura y Conceptos de Go

### Conceptos Clave Aplicados
* **Goroutines & Channels**: Como pilares para la concurrencia, comunicación y sincronización.
* **`sync.WaitGroup`**: Para orquestar la finalización ordenada de la simulación.
* **Estructura de Paquetes**: Layout `cmd/` e `internal/` para una clara separación de responsabilidades.

### Estructura de Directorios
```text
src/
├── cmd/
│   └── fisio-center/   # Proceso principal de la simulación (main.go)
└── internal/
    ├── centro/         # Lógica de negocio y sincronización del centro
    └── cliente/        # Definición y ciclo de vida de los clientes
```
