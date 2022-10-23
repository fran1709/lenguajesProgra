class Agenda(object):
    # Atributos
    objetos = [object]
    
    # Constructor
    def __init__(self):
        pass
    
    # Methods
    def agregar_objeto(self, pObject):
        self.objetos.append(pObject)
        print("Objeto Agregado!")
        
    def ver_lista_objetos(self):
        for i in range(len(self.objetos)):
            print(self.objetos[i])

    def obtener_lista(self):
        text = ""
        for i in range(len(self.objetos)):
            text += str(self.objetos[i]) + "\n"
        return text