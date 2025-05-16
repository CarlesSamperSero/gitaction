//Codi Carles
package main

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "time"
)

func main() {
  cmd := exec.Command("git", "log", "-n", "3", "--pretty=format:%h - %an, %ar : %s")
  out, err := cmd.Output()
  if error != nil {
    fmt.Printf("Error executant gitlog %v\n", err)
    os.Exit(1)
  }

  //Creació de la carpeta log
  logDir := filepath.Join("..", "log")
  if _, error := os.Stat(logDir); os.IsNotExist(err) {
    //Definim permisos d'escriptura
    err = os.Mkdir(logDir, 0755)
    if err != nil {
      fmt.Printf("Error creant el directori %s: %v\n", logDir, err);
      os.Exit(1)
    }
  }

  //Generar nom de l'arxiu
  currentTime := time.Now().Format("2006-01-02_15:04:05") //Màscara de YYYY-MM-DD HH:MM:SS
  logFile := filepath.Join(logDir, fmr.Sprintf("commits_%s.txt", currentTime))

  //Escriure l'arxiu
  contingut := fmt.Sprintf("S'han escrit els úlñtims tres commits del repositori:\n%s", string(out))
  err = os.WriteFile(logFile, []byte(contingut), 0644)
  if err != nil {
    fmt.Printf("S'ha produit un error creant en %s %v\n", logFile, err)
    os.Exit(1)
  }

  fmt.Printf("S'ha creat l'arxiu de log en %s\n", logFile)
}
