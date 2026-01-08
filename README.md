# send-promt-ai
backend microservice that acts as an asynchronous gateway between client applications and Large Language Model (LLM) providers like OpenAI


###
Componente,Capa,Ubicación,Responsabilidad
AIServicePort,Dominio,domain/ports,Define qué puede hacer un servicio de IA.
OpenAIAdapter,Infraestructura,infrastructure/adapters,Implementa la llamada real a OpenAI.
ProcessAIUseCase,Aplicación,application/usecases,Maneja la cola de trabajos y las goroutines.
AIHandler,Infraestructura,infrastructure/http/handlers,Traduce el protocolo HTTP a lógica de aplicación.
RunServer,Infraestructura,infrastructure/http,Configura y levanta el servidor net/http.
###