# Purchases

## Caráteristicas principales

1. Listas de compra sencillas
    - **Productos como texto libre**(sin IDs, solo nombres)
    - **Marcar/Desmarcar** como comprado
        * Al marcar: **precio obligatorio** (campo de texto, ej: "5.99").
        * Lugar de compra **opcional** (default:"supermercado")

2. Clonar listas
    - Crear una nueva lista a partir de una existente con todos los productos **desmarcados**
    - Ideal para listas recurrentes como compras semanales

3. Base de datos liviana
    - SQLite (sqlitedbfile.db) + GORM para ORM
    - Sin autenticación(Solo datos locales)

4. Tecnologías simplificadas
    - Backend con Go Puro(`net/http`)
    - GORM para interactuar con SQLITE
    - Validator para la validación de campos obligatorios
    - Plantillas HTML integradas (sin frameworks frontend)

5. Estructura del proyecto

```plaintext
Purchases
    |---main.go             // Punto de entrada
    |---handlers/           // Lógica de endpoints 
    |   |---list.go
    |   |---items.go
    |---models/             // Modelos de DB
    |   |---shopping_list.go
    |   |---list_item.go
    |---templates/          // Platillas HTML
    |   |---index.html      // Listar todas las listas
    |   |---list.html       // Detalle de una lista
    |---sqlitefile.db       // Base de datos
```
