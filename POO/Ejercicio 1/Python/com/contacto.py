from com.persona import Persona

class Contacto(object):
    # Atributos
    persona = Persona
    celular = str
    
    def __init__(self, pPersona):
        self.persona = pPersona
        
    def get_celular(self) -> str:
        return self.celular
    
    def __repr__(self):
        text = ""
        text += "Nombre -> " + self.persona.get_nombre_completo() + "\n"
        text += "Edad -> " + self.persona.get_edad() + "\n"
        text += "Cédula -> " + self.persona.get_cedula() + "\n"
        text += "Celular -> " + self.get_celular() + "\n"
        return text
        
class Contacto_Trabajo(Contacto):
    # Atributos
    empresa = str
    telefono = str
    correo = str
    
    # Constructor
    def __init__(self, pPersona, pEmpresa, pTelefono, pCelular, pCorreo):
        self.persona = pPersona
        self.empresa = pEmpresa
        self.telefono = pTelefono
        self.celular = pCelular
        self.correo = pCorreo
       
    def get_empresa(self)->str:
        return self.empresa
    def get_telefono(self)->str:
        return self.telefono
    def get_correo(self)->str:
        return self.correo
        
    def __repr__(self):
        text = super().__repr__()
        text += "Empresa - > " + self.get_empresa() + "\n"
        text += "Teléfono -> " + self.get_telefono() + "\n"
        text += "Correo -> " + self.get_correo() + "\n"
        return text
    
        return text
     
class Contacto_Familiar(Contacto):
    # Atributos
    parentezco = str
    
    # Construtor 
    def __init__(self, pPersona, pCelular, pParentezco):
        self.persona = pPersona
        self.celular = pCelular
        self.parentezco = pParentezco
    
    def get_parentezco(self):
        return self.parentezco
    
    def __repr__(self):
        text = super().__repr__()
        text += "Parentezco - > " + self.get_parentezco() + "\n"
        return text
    