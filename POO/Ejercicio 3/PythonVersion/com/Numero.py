class Numero:
    # Atributos
    valor = int
    
    # Constructor
    def __init__(self, pNum):
        self.valor = pNum
    
    # Métodos
    def get_valor(self) -> int:
        return self.valor
    def set_valor(self, pNum):
        self.valor = pNum
    
    def to_string(self) -> str:
        text = str
        text += str(self.get_valor())
        return text