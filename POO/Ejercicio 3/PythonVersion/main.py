from com.Solve import Solve

if __name__=="__main__":
   Calculadora = Solve("Fran's Calculator")
   operacion = "512*30/2+320-1"
   print(Calculadora.enlistar(operacion))