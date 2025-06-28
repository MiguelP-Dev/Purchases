# 🛒 Purchases - Listas de Compras

Una aplicación web sencilla para organizar listas de compras con Go, GORM y SQLite.

## ✨ Características Principales

### 📝 Listas de Compras Sencillas

- **Productos como texto libre** (sin IDs, solo nombres)
- **Marcar/Desmarcar** como comprado
- **Precio obligatorio** al marcar como comprado
- **Lugar de compra opcional** (default: "supermercado")

### 🔄 Clonar Listas

- Crear una nueva lista a partir de una existente
- Todos los productos se clonan **desmarcados**
- Ideal para listas recurrentes como compras semanales

### 💾 Base de Datos Liviana

- **SQLite** (sqlitefile.db) + **GORM** para ORM
- Sin autenticación (Solo datos locales)
- Auto-migración de esquemas

### 🛠️ Tecnologías Simplificadas

- **Backend**: Go Puro (`net/http`)
- **ORM**: GORM para interactuar con SQLite
- **Router**: Gorilla Mux
- **Plantillas**: HTML integradas con CSS inline
- **UI**: Diseño moderno con gradientes y animaciones

## 📁 Estructura del Proyecto

```plaintext
Purchases/
├── main.go                 # Punto de entrada
├── go.mod                  # Dependencias
├── go.sum                  # Checksums de dependencias
├── LICENSE                 # Licencia MIT
├── handlers/               # Lógica de endpoints
│   ├── list.go            # Handlers de listas
│   └── items.go           # Handlers de items
├── models/                 # Modelos de DB
│   ├── shopping_list.go   # Modelo de lista
│   └── list_item.go       # Modelo de item
├── templates/              # Plantillas HTML
│   ├── index.html         # Listar todas las listas
│   ├── list.html          # Detalle de una lista
│   ├── create_list.html   # Crear nueva lista
│   ├── edit_list.html     # Editar nombre de lista
│   ├── mark_purchased.html # Marcar como comprado
│   ├── edit_item.html     # Editar producto
│   ├── favicon.svg        # Favicon con carrito de compras
│   └── functions.go       # Funciones de plantilla
└── sqlitefile.db          # Base de datos SQLite
```

## 🚀 Instalación y Uso

### Prerrequisitos

- Go 1.24.4 o superior

### 1. Clonar el Repositorio

```bash
git clone <tu-repositorio>
cd Purchases
```

### 2. Instalar Dependencias

```bash
go mod tidy
```

### 3. Ejecutar la Aplicación

```bash
go run main.go
```

### 4. Acceder a la Aplicación

Abre tu navegador y ve a: `http://localhost:8080`

## 📋 Funcionalidades

### Crear Lista

1. Haz clic en "➕ Crear Nueva Lista"
2. Ingresa un nombre descriptivo
3. Haz clic en "Crear Lista"

### Agregar Productos

1. Ve a una lista específica
2. Usa el formulario "➕ Agregar Producto"
3. Ingresa el nombre del producto
4. Haz clic en "Agregar Producto"

### Marcar como Comprado

1. Haz clic en "Comprar" junto al producto
2. Ingresa el **precio** (obligatorio)
3. Ingresa el **lugar** (opcional)
4. Haz clic en "✅ Confirmar Compra"

### Editar Producto Comprado

1. Haz clic en "✏️ Editar" junto al producto comprado
2. Modifica el **precio** y/o **lugar**
3. Haz clic en "💾 Guardar Cambios"

### Editar Nombre de Lista

1. Haz clic en "✏️ Editar Nombre" en el encabezado de la lista
2. Modifica el **nombre** de la lista
3. Haz clic en "💾 Guardar Cambios"

### Alternar Estado

- Usa el checkbox para marcar/desmarcar rápidamente
- Los productos comprados se muestran tachados

### Clonar Lista

1. Haz clic en "Clonar" en la tarjeta de la lista
2. Se crea una copia con todos los productos desmarcados
3. El título se agrega "(copia)"

### Eliminar

- **Lista**: Haz clic en "Eliminar" en la tarjeta de la lista
- **Producto**: Haz clic en "Eliminar" junto al producto

## 🎨 Características de la UI

### Diseño Moderno

- Gradientes de colores atractivos
- Animaciones suaves en hover
- Diseño responsive
- Iconos emoji para mejor UX
- Favicon con carrito de compras

### Indicadores Visuales

- Barra de progreso en cada lista
- Productos comprados tachados
- Colores diferenciados por acción
- Estados visuales claros

### Navegación Intuitiva

- Botones de "Volver" consistentes
- Confirmaciones para eliminaciones
- Formularios con validación
- Mensajes de ayuda contextuales

## 🔧 Configuración

### Puerto del Servidor

Por defecto la aplicación corre en el puerto 8080. Para cambiar:

```go
// En main.go, línea 67
port := ":8080" // Cambia a tu puerto preferido
```

### Base de Datos

La aplicación usa SQLite por defecto. El archivo se crea automáticamente en:

```plaintext
sqlitefile.db
```

## 🐛 Solución de Problemas

### Error de Dependencias

```bash
go mod tidy
go mod download
```

### Error de Base de Datos

- Elimina `sqlitefile.db` si existe
- La aplicación creará una nueva base de datos

### Error de Plantillas

- Verifica que todos los archivos HTML estén en `templates/`
- Asegúrate de que los nombres de archivo coincidan

### Puerto en Uso

```bash
# Encuentra el proceso usando el puerto 8080
lsof -i :8080
# Mata el proceso
kill -9 <PID>
```

## 📝 API Endpoints

### Listas

- `GET /` - Página principal
- `GET /create` - Formulario crear lista
- `POST /create` - Crear lista
- `GET /list/{id}` - Ver lista específica
- `GET /list/{id}/edit` - Formulario editar lista
- `POST /list/{id}/edit` - Editar lista
- `POST /list/{id}/delete` - Eliminar lista
- `POST /list/{id}/clone` - Clonar lista
- `POST /list/{id}/items` - Agregar item

### Items

- `POST /items/{id}/delete` - Eliminar item
- `POST /items/{id}/toggle` - Alternar estado
- `GET /items/{id}/mark-purchased` - Formulario marcar comprado
- `POST /items/{id}/mark-purchased` - Marcar como comprado
- `GET /items/{id}/edit` - Formulario editar item
- `POST /items/{id}/edit` - Editar item
- `POST /items/{id}/update` - Actualizar item

## 🛠️ Desarrollo

### Compilar

```bash
go build -o purchases main.go
```

### Ejecutar Binario

```bash
./purchases
```

### Estructura de Base de Datos

```sql
-- ShoppingList
- ID (uint, primary key)
- Title (string, not null)
- CreatedAt (time)
- UpdatedAt (time)

-- ListItem
- ID (uint, primary key)
- ShoppingListID (uint, foreign key)
- ProductName (string, not null)
- IsPurchased (bool, default false)
- Price (float64, default 0.0)
- Place (string, default 'supermercado')
- CreatedAt (time)
- UpdatedAt (time)
```

## 🤝 Contribuir

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 🙏 Agradecimientos

- [Gorilla Mux](https://github.com/gorilla/mux) - Router HTTP
- [GORM](https://gorm.io/) - ORM para Go
- [SQLite](https://www.sqlite.org/) - Base de datos ligera
