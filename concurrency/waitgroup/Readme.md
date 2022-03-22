##  Waitgroup

Documentacion oficial [WaitGroup](https://translate.google.com.mx/?hl=es)

1. Create WaitGroup
2. Pasamos el WaitGroup como parametro a una funcion 
3. Sumamos 1 antes de la ejecucion con wg.Add(1) 
4. Restamos 1 dentro de la funcion con wg.Done()
4. Usanos wg.Wait() para bloquear la ejecucion y que terminen de ejecutarse todas las funciones del WaitGroup