class Producto(object):
    #Atirbutos
    nombre = str
    precio = int
    cantidad = int
    
    # Constructor
    def __init__(self, pName, pPrecio, pCantidad):
        self.nombre = pName
        self.precio = pPrecio
        self.cantidad = pCantidad
        
    # Methods
    def get_nombre(self)->str:
        return self.nombre
    def get_precio(self)->str:
        return str(self.precio)
    def get_cantidad(self)->str:
        return str(self.cantidad)
    
    def __repr__(self):
        text = ""
        text += "Nombre -> " + self.get_nombre() + "\n"
        text += "Precio -> " + self.get_precio() + "\n"
        text += "Cantidad -> " + self.get_cantidad() + "\n"
        return text