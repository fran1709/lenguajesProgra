from tkinter import Tk
from com.contacto import Contacto_Trabajo
from com.contacto import Contacto_Familiar
from com.producto import Producto
from gui.gui import Window
from com.persona import Persona
from com.agenda import Agenda
 

if __name__ == '__main__':
    #Instanciando
    agenda_fran = Agenda()
    
    # Agregando
    agenda_fran.agregar_objeto(Contacto_Familiar(Persona("Francisco","Ovares","Rojas",25,207710202),"8514-6189","Esposo"))
    agenda_fran.agregar_objeto(Contacto_Familiar(Persona("Valeria","Molina","Portuguez",21,207710205),"5674-6344","Esposa"))
    agenda_fran.agregar_objeto(Contacto_Trabajo(Persona("Walter","Lazo","González",19,"208410988"),"NASA","1-23123-2313","86367168","eltec@estudiantec.cr"))
    agenda_fran.agregar_objeto(Producto("Pizza Grande", 12000, 3))
    
    # Mostrando en consola
    agenda_fran.ver_lista_objetos()
    # Mostando en GUI
    root = Tk()
    miVentana = Window(root, agenda_fran.obtener_lista())
    root.mainloop() 