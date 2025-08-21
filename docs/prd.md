# Prompt Requirement Document (PRD): Vitykunja
**Versión:** 1.1

---

### **G1: Guideline (Directriz) - El "Porqué"**

#### **Visión y Objetivos**
* **Visión:** Crear **Vitykunja**, el primer gestor de tareas "AI-native" y autoalojado, que ofrezca una experiencia de usuario fluida a través de un NLP superior en español y una integración perfecta y granular con el ecosistema de Google Calendar.
* **Objetivos:**
    * **MVP:** Lograr una sincronización bidireccional y en tiempo real entre un proyecto de Vitykunja y un calendario específico de Google.
    * **KPI #1:** El 95% de las tareas creadas/actualizadas/borradas en Vitykunja se reflejan en Google Calendar en menos de 10 segundos, y viceversa.
    * **KPI #2:** Implementar un endpoint `/mcp` que describa los modelos de datos (`Task`, `Project`) y las acciones (`createTask`, `findTasks`) disponibles para ser consumido por IAs externas.

#### **Requisitos del Sistema**
* **Autenticación Google:** Los usuarios deben poder conectar su cuenta de Google de forma segura mediante un flujo OAuth 2.0.
* **Mapeo Proyecto-Calendario:** En la configuración de cada proyecto de Vitykunja, el usuario debe poder seleccionar uno de sus calendarios de Google para la sincronización.
* **Sincronización Bidireccional:**
    * Las tareas con fecha en Vitykunja (creadas, actualizadas, eliminadas) deben sincronizarse como eventos en el calendario de Google mapeado.
    * Los eventos en el calendario de Google mapeado (creados, actualizados, eliminados) deben sincronizarse como tareas en Vitykunja.
* **Motor de NLP en Español:** El parseo de fechas/horas se mantiene como una funcionalidad clave para la creación rápida de tareas.
* **Model Context Protocol (MCP):** Un endpoint público `GET /mcp` que no requiere autenticación y devuelve un archivo `mcp.yaml`.

#### **Stack Tecnológico y Decisiones de Arquitectura**
* **Backend (Go):**
    * Se usará la librería oficial `google.golang.org/api/calendar/v3` para interactuar con la API de Google.
    * Las credenciales de usuario (tokens de acceso/refresco) se almacenarán en la base de datos de forma encriptada.
    * Se implementará un servicio de `background worker` (usando Go routines y channels) para gestionar la cola de sincronización y evitar bloqueos en la API principal.
* **Frontend (Vue.js):**
    * Crear una nueva vista de "Integraciones" en la configuración del usuario para gestionar la conexión con Google.
    * Añadir un componente en la configuración de cada proyecto para seleccionar el calendario de destino.

---

### **G2: Guidance (Guía) - El "Cómo"**

#### **Metodología y Ejemplos de Prompts para la IA**

* **Para implementar el flujo OAuth 2.0 en Go:**
    > "Actúa como un experto en Go y seguridad. Necesito implementar un flujo de autorización OAuth 2.0 para la API de Google Calendar en el backend de Vitykunja. Crea un nuevo servicio `google_auth_service.go`. Debe contener tres handlers: `HandleLogin` (que redirige al usuario a la página de consentimiento de Google), `HandleCallback` (que recibe el código de autorización, lo intercambia por un token y guarda el token de refresco encriptado en la base de datos asociado al usuario) y `RefreshToken` (una función que usa el token de refresco para obtener un nuevo token de acceso cuando expire)."

* **Para crear el servicio de sincronización bidireccional:**
    > "Genera la estructura para un servicio de sincronización bidireccional en Go llamado `sync_service.go`. Debe tener dos funciones principales: `SyncToGoogle(task models.Task)` y `SyncFromGoogle(event calendar.Event)`. La `SyncToGoogle` se llamará cuando una tarea se cree/actualice en Vitykunja. La `SyncFromGoogle` se ejecutará en un worker que periódicamente consulta los calendarios mapeados en busca de cambios usando un `syncToken` de la API de Google. Diseña la lógica para evitar bucles de sincronización infinitos, por ejemplo, guardando el ID del evento de Google en la tarea de Vitykunja y viceversa."

* **Para definir el Model Context Protocol (MCP):**
    > "Crea el contenido para un archivo `mcp.yaml`. Este archivo debe definir el contexto para una IA que quiera interactuar con Vitykunja. Debe incluir una sección `models` que describa la estructura de una `Task` (con campos como id, title, dueDate, projectId) y un `Project`. También debe incluir una sección `actions` que defina cómo interactuar. Por ejemplo, la acción `createTask` debe especificar los parámetros requeridos (`title`, `projectId`) y opcionales (`dueDate`), y la ruta de la API a la que llamar (`POST /api/v1/tasks`). Este YAML será la única fuente de verdad para la IA."

---

### **G3: Guardrails (Barandillas) - Los "Límites"**

#### **Estándares y Reglas de Calidad**
* **Seguridad:** **Prioridad máxima.** Los tokens de la API de Google deben estar siempre encriptados en la base de datos. El `client_secret` de la aplicación de Google no debe estar hardcodeado, se cargará desde variables de entorno.
* **Privacidad de Datos:** La aplicación solo debe solicitar los permisos estrictamente necesarios para la sincronización (`calendar.events` y `calendar.readonly`). La política de privacidad debe ser clara al respecto.
* **Rendimiento y Límites de API:** El `background worker` de sincronización debe respetar los límites de uso de la API de Google. Implementar un mecanismo de `exponential backoff` para reintentar las llamadas fallidas.
* **Consistencia de Datos:** La sincronización debe ser transaccional siempre que sea posible. Si una actualización falla a mitad de camino, se debe poder revertir o marcar para una resincronización.
* **MCP Estándar:** El archivo `mcp.yaml` debe seguir una estructura clara y consistente, versionada junto con la API. Cualquier cambio en la API que afecte a la IA debe reflejarse primero en el MCP.
