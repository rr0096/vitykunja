# Vitykunja: El Gestor de Tareas AI-Native

**Vitykunja** es un gestor de tareas autoalojado de nueva generación, construido sobre la sólida base de [Vikunja](https://vikunja.io/). Diseñado para "Vibe Coders" y usuarios avanzados, Vitykunja integra un **NLP en español** de alta precisión, una **sincronización bidireccional por proyecto con Google Calendar** y un **Model Context Protocol (MCP)** para una interacción nativa con IA.

## ✨ Características Principales

* **Sincronización Profunda con Google Calendar:** Olvídate de CalDAV. Vitykunja se integra directamente con la API de Google Calendar. Mapea cada uno de tus proyectos a un calendario específico y disfruta de una sincronización bidireccional y en tiempo real.
* **Creación de Tareas con NLP en Español:** Escribe de forma natural. `"Reunión de equipo próximo viernes a las 10am #trabajo"` creará la tarea, la asignará al proyecto correcto y la pondrá en el calendario adecuado.
* **Preparado para la IA con MCP:** Vitykunja es pionero en incluir un Model Context Protocol. Expone una definición clara y estructurada de sus capacidades, permitiendo que agentes de IA externos puedan leer, entender e interactuar con tus tareas de forma segura y predecible.
* **100% Autoalojado y de Código Abierto:** Control total sobre tus datos, tu flujo de trabajo y tu privacidad.

## 🚀 Stack Tecnológico

* **Backend:** Go
* **Frontend:** Vue.js
* **Sincronización:** Google Calendar API v3 (OAuth 2.0)
* **Protocolo IA:** Model Context Protocol (MCP) vía endpoint `/mcp`
* **Especificación MCP:** `docs/mcp.yaml` — especificación de modelos y acciones para agentes
* **Deployment:** Docker

## 🛠️ Cómo Empezar

### Prerrequisitos
* [Docker](https://www.docker.com/get-started)
* [Docker Compose](https://docs.docker.com/compose/install/)
* Credenciales de la API de Google (Client ID y Client Secret)

### Instalación
1.  **Clona el repositorio:**
    ```bash
    git clone [https://github.com/tu-usuario/vitykunja.git](https://github.com/tu-usuario/vitykunja.git)
    cd vitykunja
    ```

2.  **Crea un proyecto en Google Cloud Console:** Habilita la API de Google Calendar y obtén tus credenciales OAuth 2.0. Asegúrate de añadir la URL de callback correcta (ej. `http://localhost:8080/api/v1/auth/google/callback`) a las URIs de redirección autorizadas.

3.  **Configura tu `.env`:**
    Copia el archivo de ejemplo y rellena tus credenciales.
    ```bash
    cp .env.example .env
    ```
    Edita el `.env` y añade tu `GOOGLE_CLIENT_ID` y `GOOGLE_CLIENT_SECRET`.

4.  **Levanta los contenedores:**
    ```bash
    docker-compose up --build -d
    ```
5.  **¡Accede y Conecta!**
    Abre tu navegador en `http://localhost:8080`. Ve a la configuración de tu perfil, busca la sección de integraciones y conecta tu cuenta de Google.

## 🗺️ Hoja de Ruta (Roadmap)

* **✅ Fase 1: MVP - Integración con Google y MCP** - ¡En progreso!
* **▶️ Fase 2: NLP Mejorado** - Reconocimiento de proyectos, etiquetas y prioridades en el texto.
* **⏹️ Fase 3: Ecosistema de IA y Agentes** - Desarrollo de agentes de ejemplo y ampliación del MCP.
