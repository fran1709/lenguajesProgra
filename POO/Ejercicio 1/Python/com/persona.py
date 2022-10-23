class Persona(object):
    # Atributos.
    nombre = str
    apellido1 = str
    apellido2 = str
    edad = int
    cedula = int
    
    # Constructor.
    def __init__(self,pNombre,pApellido1,pApellido2,pEdad,pCedula):
        self.nombre = pNombre
        self.apellido1 = pApellido1
        self.apellido2 = pApellido2
        self.edad = pEdad
        self.cedula = pCedula
        
    # GETS & SETS
    def get_nombre(self) -> str:
        return self.nombre
    def get_apellido1(self) -> str:
        return self.apellido1
    def get_apellido2(self) -> str:
        return self.apellido2
    def get_edad(self) -> str:
        return str(self.edad)
    def get_cedula(self) -> str:
        return str(self.cedula)
    def get_nombre_completo(self):
        return (self.get_nombre() + " " + self.get_apellido1() + " " + self.get_apellido2())
    