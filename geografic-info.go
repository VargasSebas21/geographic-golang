package main

import (
    "fmt"
    "strings"
)

// **Estructura del grafo**

type Grafo struct {
    nodos []Nodo
}

// **Nodo**

type Nodo interface {
    Exportar(formato string) string
}

// **Ciudad**

type Ciudad struct {
    nombre string
    poblacion int
}

// **Funcion Exportar() de la clase Ciudad**

func (c *Ciudad) Exportar(formato string) string {
    // **Validamos el formato**
    if formato != "xml" && formato != "json" {
        panic(fmt.Sprintf("Formato no soportado: %s", formato))
    }

    // **Generamos el XML o JSON**
    if formato == "xml" {
        return fmt.Sprintf(
            "<nodo tipo=\"ciudad\"> <nombre>%s</nombre> <poblacion>%d</poblacion> </nodo>",
            c.nombre,
            c.poblacion,
        )
    } else {
        return fmt.Sprintf(
            `{ "tipo": "ciudad", "nombre": "%s", "poblacion": %d }`,
            c.nombre,
            c.poblacion,
        )
    }
}

// **Industria**

type Industria struct {
    tipo string
    empleados int
}

// **Funcion Exportar() de la clase Industria**

func (i *Industria) Exportar(formato string) string {
    // **Validamos el formato**
    if formato != "xml" && formato != "json" {
        panic(fmt.Sprintf("Formato no soportado: %s", formato))
    }

    // **Generamos el XML o JSON**
    if formato == "xml" {
        return fmt.Sprintf(
            "<nodo tipo=\"industria\"> <tipo>%s</tipo> <empleados>%d</empleados> </nodo>",
            i.tipo,
            i.empleados,
        )
    } else {
        return fmt.Sprintf(
            `{ "tipo": "industria", "tipo": "%s", "empleados": %d }`,
            i.tipo,
            i.empleados,
        )
    }
}

// **LugarTuristico**

type LugarTuristico struct {
    nombre string
    tipo string
}

// **Funcion Exportar() de la clase LugarTuristico**

func (l *LugarTuristico) Exportar(formato string) string {
    // **Validamos el formato**
    if formato != "xml" && formato != "json" {
        panic(fmt.Sprintf("Formato no soportado: %s", formato))
    }

    // **Generamos el XML o JSON**
    if formato == "xml" {
        return fmt.Sprintf(
            "<nodo tipo=\"lugar_turistico\"> <nombre>%s</nombre> <tipo>%s</tipo> </nodo>",
            l.nombre,
            l.tipo,
        )
    } else {
        return fmt.Sprintf(
            `{ "tipo": "lugar_turistico", "nombre": "%s", "tipo": "%s" }`,
            l.nombre,
            l.tipo,
        )
    }
}

// **Funcion main()**

func main() {
    // **Creamos un grafo con algunos nodos**
    grafo := Grafo{}
    grafo.nodos = []Nodo{
        &Ciudad{nombre: "Bogotá", poblacion: 10_000_000},
        &Industria{tipo: "Tecnología", empleados: 100_000},
        &LugarTuristico{nombre: "Catedral Primada de Bogotá", tipo: "Arquitectónico"},
    }

    // **Exportamos el grafo a formato XML**
    xml := grafo.Exportar("xml")

    // **Imprimimos el XML**
    fmt.Println(xml)
}