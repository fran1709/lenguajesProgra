import re

class Solve:
    # Atributos
    nombre = str
    
    # Constructor
    def __init__(self, pName):
        self.nombre = pName
    
    # Métodos
    def enlistar(self, pOperation):
        datos = re.split('[/ * + -]', pOperation)
        return datos
        