# golangci-lint para repo mono módulo

### Descripción

Ejemplo de cómo usar el linter para repositorios donde existe sólo un ``go.mod``.

### ¿Qué contiene?

- Archivo de configuración del linter (`.golangci.yml`).
- Github workflow para correr el linter y tests en el CI/CD.
- Directorio con dos git hooks:
  - ``pre-commit``: corre el linter sólo si se modifican archivos de Go.
  - ``pre-push``: corre los unit tests antes de pushear en caso de que se commiteen archivos de Go.
- Makefile que se encarga de instalar las dependencias (de ser necesario), correr los tests y linter e instalar los git hooks.

## Primeros pasos

Ejecutar el comando `make install-hooks` para instalar los git hooks localmente.

### Dependencias

- [`golangci-lint`](https://golangci-lint.run/)
- [`husky`](https://github.com/automation-co/husky)
