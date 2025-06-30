# 🛒 Listas de Compras

Una aplicación web simple y eficiente para gestionar listas de compras, desarrollada en Go con una interfaz moderna y responsive.

## ✨ Características

- **Gestión de Listas**: Crear, editar, eliminar y clonar listas de compras
- **Gestión de Productos**: Agregar, marcar como comprado, editar y eliminar productos
- **Seguimiento de Compras**: Registrar precio y lugar de compra para cada producto
- **Progreso Visual**: Barra de progreso que muestra el avance de las compras
- **Interfaz Responsive**: Diseño moderno que funciona en dispositivos móviles y desktop
- **Base de Datos SQLite**: Almacenamiento local simple y eficiente

## 🚀 Tecnologías Utilizadas

- **Backend**: Go 1.24.4
- **Framework Web**: Gorilla Mux
- **Base de Datos**: SQLite con GORM
- **Frontend**: HTML5, CSS3, JavaScript vanilla
- **Templates**: Go HTML Templates

## 📋 Requisitos

- Go 1.24.4 o superior
- Git

## 🛠️ Instalación

1. **Clonar el repositorio**:

   ```bash
   git clone https://github.com/MiguelP-Dev/Purchases.git
   cd Purchases
   ```

2. **Instalar dependencias**:

   ```bash
   go mod download
   ```

3. **Ejecutar la aplicación**:

   ```bash
   go run main.go
   ```

4. **Abrir en el navegador**:

   ```
   http://localhost:8080
   ```

## 📁 Estructura del Proyecto

```
Purchases/
├── main.go                 # Punto de entrada de la aplicación
├── go.mod                  # Dependencias de Go
├── go.sum                  # Checksums de dependencias
├── sqlitefile.db          # Base de datos SQLite (generada automáticamente)
├── .gitignore             # Archivos ignorados por Git
├── LICENSE                # Licencia del proyecto
├── README.md              # Este archivo
├── handlers/              # Controladores HTTP
│   ├── list.go           # Handlers para listas de compras
│   └── items.go          # Handlers para productos
├── models/               # Modelos de datos
│   ├── shopping_list.go  # Modelo de lista de compras
│   └── list_item.go      # Modelo de producto
└── templates/            # Plantillas HTML y assets
    ├── index.html        # Página principal
    ├── list.html         # Vista de lista individual
    ├── create_list.html  # Formulario de creación
    ├── edit_list.html    # Formulario de edición
    ├── edit_item.html    # Edición de productos
    ├── mark_purchased.html # Marcado como comprado
    ├── styles.css        # Estilos CSS
    ├── favicon.svg       # Icono de la aplicación
    └── functions.go      # Funciones de template
```

## 🗄️ Modelos de Datos

### ShoppingList

- `ID`: Identificador único
- `Title`: Título de la lista
- `CreatedAt`: Fecha de creación
- `UpdatedAt`: Fecha de última modificación
- `Items`: Relación con los productos (1:N)

### ListItem

- `ID`: Identificador único
- `ShoppingListID`: ID de la lista padre
- `ProductName`: Nombre del producto
- `IsPurchased`: Estado de compra
- `Price`: Precio del producto
- `Place`: Lugar de compra
- `CreatedAt`: Fecha de creación
- `UpdatedAt`: Fecha de última modificación

## 🔗 Rutas de la API

### Listas de Compras

- `GET /` - Página principal con todas las listas
- `GET /create` - Formulario de creación de lista
- `POST /create` - Crear nueva lista
- `GET /list/{id}` - Ver lista específica
- `GET /list/{id}/edit` - Formulario de edición
- `POST /list/{id}/edit` - Actualizar lista
- `POST /list/{id}/delete` - Eliminar lista
- `POST /list/{id}/clone` - Clonar lista
- `POST /list/{id}/items` - Agregar producto

### Productos

- `POST /items/{id}/delete` - Eliminar producto
- `POST /items/{id}/toggle` - Cambiar estado de compra
- `GET /items/{id}/mark-purchased` - Formulario de precio
- `POST /items/{id}/mark-purchased` - Marcar como comprado
- `GET /items/{id}/edit` - Formulario de edición
- `POST /items/{id}/edit` - Actualizar producto

## 🎨 Características de la Interfaz

### Diseño Responsive

- Adaptable a dispositivos móviles y desktop
- Navegación intuitiva con botones de acción claros
- Formularios con validación en tiempo real

### Funcionalidades Principales

- **Crear Listas**: Formulario simple para crear nuevas listas
- **Agregar Productos**: Campo de texto para agregar productos rápidamente
- **Marcar como Comprado**: Checkbox que redirige a formulario de precio
- **Editar Productos**: Modificar precio y lugar de compra
- **Progreso Visual**: Barra de progreso que muestra el avance
- **Clonar Listas**: Crear copias de listas existentes

### Estados de Productos

- **No Comprado**: Producto pendiente de compra
- **Comprado**: Producto con precio y lugar registrados

## 🔧 Configuración

La aplicación se ejecuta por defecto en el puerto 8080. Para cambiar el puerto, modifica la variable `port` en `main.go`:

```go
port := ":8080" // Cambiar por el puerto deseado
```

## 📊 Base de Datos

La aplicación utiliza SQLite como base de datos local. El archivo `sqlitefile.db` se crea automáticamente al ejecutar la aplicación por primera vez.

### Migración Automática

GORM maneja automáticamente la creación de tablas y migraciones de esquema.

## 🚀 Despliegue

### Desarrollo Local

```bash
go run main.go
```

### Compilación para Producción

```bash
go build -o purchases main.go
./purchases
```

### Docker (Opcional)

```dockerfile
FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o purchases main.go
EXPOSE 8080
CMD ["./purchases"]
```

## 🤝 Contribuir

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📝 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 👨‍💻 Autor

**Miguel P.** - [GitHub](https://github.com/MiguelP-Dev)

## 🙏 Agradecimientos

- [Gorilla Mux](https://github.com/gorilla/mux) - Router HTTP
- [GORM](https://gorm.io/) - ORM para Go
- [SQLite](https://www.sqlite.org/) - Base de datos ligera

---

⭐ Si te gusta este proyecto, ¡dale una estrella en GitHub!
