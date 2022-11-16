class Operador:
    # Atributos
    signo = str
    
    # Constructor
    def __init__(self, pOperador) -> None:
        self.signo = pOperador
        
    # Métodos
    def get_signo(self) -> str:
        return self.signo
    def set_signo(self, pSigno):
        self.signo = pSigno
        
    def to_string(self) -> str:
        text = str
        text += self.get_signo()
        return text