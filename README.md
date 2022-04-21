## Discovery Simulation (T1120, T1057, T1518)

La simulación de descubrimiento representa un malware que ejecuta tres técnicas de la táctica Discovery en MITRE ATT&CK, realizando un reconocimiento de la información del sistema que se utiliza posteriormente para explotar aún más el sistema:

Esto incluye, en este orden:
- Descubre los procesos en ejecución en el sistema.
- Descubre los periféricos presentes en el sistema.
- Descubre el software instalado en el sistema con sus respectivas versiones.

Sandbox Analysis 

| Sandbox |
| ------- |
| [Hybrid-Analysis](https://www.hybrid-analysis.com/sample/c8fcd8419bf11385bdddc9cfd8017226493365ff97d2232f9283fbe6309830bc/61dff860d9a3de1d1f04a1fb) |
| [Intezer Analyze](https://analyze.intezer.com/analyses/0e7f5d02-1f69-43a5-b6b3-65c3cffcd21d) |

## UAC Bypass Simulation

La simulación UAC Bypass consiste en un malware que utiliza la utilidad fodhelper.exe disponible en Windows 10 para lograr una escalada de privilegios locales mediante la creación de una estructura de registro para ejecutar comandos arbitrarios con privilegios de administrador:

Esto incluye, en este orden:
- Crea una nueva estructura de registro en `HKCU:\Software\Classes\ms-settings\` e inicie `notepad.exe` con privilegios de administrador evadiendo UAC.

Sandbox Analysis

| Sandbox |
| ------- |
| [Hybrid-Analysis](https://www.hybrid-analysis.com/sample/98ee778d81174276c74ef2039163b48479b9b1d798770ea434d8d54cb35390b0) |
| [Intezer Analyze](https://analyze.intezer.com/analyses/5a60925a-7195-4b4c-8f23-5db9dfbbea5a) |


## Registry Run Key Persistence Simulation

Esta es una simulación de técnicas de persistencia que utiliza claves Run del registro para lograr la persistencia de cargas útiles arbitrarias. 
Estas claves incluyen:`HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`, `EY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce`, `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServices`, `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce`.

Esto incluye, en este orden:
- Agregua un valor en la clave de registro en `HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run` para ejecutar un payload de prueba incrustado en el binario.
- Elimina el valor de `HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run` para devolverlo a su estado original para una simulación segura.

| Sandbox |
| ------- |
| [Hybrid-Analysis](https://www.hybrid-analysis.com/sample/353aa45090090f298af8b1d7135b33ea03c7b5b431c31367e9468366aff227b2) |
| [Intezer Analyze](https://analyze.intezer.com/analyses/9623c141-4927-43a3-acb9-38c65b6c9a5e) |

## Uso

- Requiere Go 1.17+, GNU make.

## Windows

Ransomware Simulation
```
$ make ransomware
$ ransomware.exe
```

Discovery Simulation
```
$ make discovery
$ discovery.exe
```

UAC Bypass Simulation
```
$ make uac_bypass
$ uac_bypass.exe
```

Registry Run Key Simulation
```
$ make registry_run
$ runkeyregistry.exe
```

## Linux/bash
```
$ GOOS=windows make ransomware # and so on
```
