from tkinter import *
import tkinter as tk

class Window(tk.Frame):
    def __init__(self, pMaster, pText):
        # Modificación de Frame
        super().__init__(pMaster)
        self.pack(expand=True, fill='both')
        pMaster.title("Contactos de Agenda")
        
        "Tablas"
        self.tabla_agenda = tk.Text(pMaster)
        self.tabla_agenda.insert(INSERT, pText)
        self.tabla_agenda.pack()
        
        # Instancia la ventana
        self.master = pMaster