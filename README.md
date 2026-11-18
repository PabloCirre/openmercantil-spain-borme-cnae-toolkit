# OpenMercantil Spain BORME, CIF/NIF & CNAE B2B Due Diligence Toolkit (Go & Open Data · 2026)

[![Go Reference](https://pkg.go.dev/badge/github.com/PabloCirre/openmercantil-spain-borme-cnae-toolkit.svg)](https://pkg.go.dev/github.com/PabloCirre/openmercantil-spain-borme-cnae-toolkit@v1.0.0)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Data License: CC BY 4.0](https://img.shields.io/badge/Open_Data-CC_BY_4.0-darkred.svg)](https://openmercantil.es/fuentes)

Open-source Go validation package (`bormecnae`) and quantitative reference tables for **Spanish corporate tax identification (`CIF / NIF` under Royal Decree 1065/2007)**, **B2B commercial solvency due diligence (Law 3/2004)**, **Verifactu anti-fraud billing compliance (RD 1007/2023 & RD 254/2025)**, **NACE Rev. 2 / CNAE sectoral distribution**, and **Spanish public procurement thresholds (`PLACSP` & `BDNS`)**.

![Spanish Mercantile Census CNAE & B2B Due Diligence Thresholds 2026](https://pablocirre.github.io/openmercantil-spain-borme-cnae-toolkit/openmercantil-b2b-due-diligence-2026-2x.png)

---

## 1. Live Public Registry & B2B Tools (`OpenMercantil`)

Mathematical checksum validation confirms that a 9-character Spanish `CIF` obeys the AEAT modulus-10 algorithm, **not** that the company is active or solvent. Cross-verify every entity against official Spanish public bulletins structured under CC BY 4.0 at **[OpenMercantil](https://openmercantil.es/)**:

| Verification Layer | Official Primary Source | OpenMercantil Public Tool (Free / No Login) |
| :--- | :--- | :--- |
| **Corporate Tax ID (`CIF / NIF`) & Legal Form** | AEAT (`RD 1065/2007`) · BORME · EU VIES | [Consultar y comprobar datos CIF de empresa gratis](https://openmercantil.es/como-comprobar-cif-empresa) |
| **B2B Solvency & 30/60-Day Trade Credit Limit** | BORME + PLACSP + BDNS + Ley 3/2004 | [Comprobar solvencia de una empresa y riesgo B2B](https://openmercantil.es/comprobar-solvencia-empresa) |
| **Verifactu & Mandatory B2B E-Invoicing Deadlines** | AEAT `RD 1007/2023`, `RD 254/2025` & Ley 18/2022 | [Calculadora Verifactu y Factura Electrónica por CNAE](https://openmercantil.es/verifactu-empresas-cnae) |
| **Newly Incorporated Companies Radar (52 Provinces)** | Sección Primera del BORME (AEBOE) | [Radar diario de nuevas empresas constituidas en España](https://openmercantil.es/nuevas-empresas-hoy) |
| **CNAE Sectoral Census (`103` Active Divisions)** | INE CNAE-2009 / CNAE-2025 (`RD 10/2025`) | [Directorio de sectores CNAE](https://openmercantil.es/sectores) · [CNAE 62 Informática](https://openmercantil.es/cnae/62) |
| **Public Tenders (`PLACSP`) & Grants (`BDNS`)** | Ministerio de Hacienda (`CODICE` / `IGAE`) | [Guía y buscador PLACSP](https://openmercantil.es/glosario/placsp) · [Qué es la BDNS](https://openmercantil.es/que-es-la-bdns) |
| **International KYB & EU VAT Reference (English)** | Mercantile Registry + EU VIES (`ES+CIF`) | [Spanish Company Register & CIF Search Guide](https://openmercantil.es/spanish-company-register) |

---

## 2. Spanish Corporate CIF Prefix Matrix (`RD 1065/2007`)

| Prefix Letter | Spanish Legal Entity Type | Control Character Rule |
| :---: | :--- | :--- |
| **A** | Sociedad Anónima (`S.A.`) | Numeric digit (`0–9`) |
| **B** | Sociedad de Responsabilidad Limitada (`S.L.` / `S.L.U.`) | Numeric digit (`0–9`) |
| **C / D** | Sociedad Colectiva (`C`) / Comanditaria (`D`) | Digit or Control Letter (`JABCDEFGHI`) |
| **E / H** | Comunidad de Bienes (`E`) / Propiedad Horizontal (`H`) | Numeric digit (`0–9`) |
| **F / G / J** | Cooperativa (`F`), Asociación/Fundación (`G`), Sociedad Civil (`J`) | Digit or Control Letter |
| **U / V** | Unión Temporal de Empresas (`U.T.E.`) / Agrupación Interés Económico (`V`) | Numeric (`U`) / Either (`V`) |
| **N / W** | Entidad Extranjera (`N`) / Establecimiento Permanente No Residente (`W`) | Control Letter (`A–J`) |
| **P / Q / R / S** | Corporación Local (`P`), Organismo Público (`Q`), Administración (`S`) | Control Letter (`A–J`) |

---

## 3. Go Installation & Usage

```bash
go get github.com/PabloCirre/openmercantil-spain-borme-cnae-toolkit@v1.0.0
```

```go
package main

import (
    "fmt"
    bormecnae "github.com/PabloCirre/openmercantil-spain-borme-cnae-toolkit"
)

func main() {
    valid, form, err := bormecnae.ValidateCIF("A28017895")
    fmt.Println(valid, form.Description, err)
    vies, _ := bormecnae.FormatEUVIES("A28017895")
    fmt.Println("EU VIES:", vies)
}
```

---

## 4. Public Procurement (`PLACSP`) & Trade Credit Statutory Thresholds (2026)

| Statutory Framework | Legal Norm (BOE) | Key Quantitative Threshold | Operational Requirement |
| :--- | :--- | :--- | :--- |
| **Minor Contract (*Contrato Menor*)** | Art. 118 Ley 9/2017 LCSP | `< 15.000 €` services / `< 40.000 €` works (ex-VAT) | Quarterly publication on [PLACSP](https://openmercantil.es/glosario/placsp) |
| **Simplified Open Tender** | Art. 159 Ley 9/2017 LCSP | `<= 143.000 €` services / `<= 2.000.000 €` works | Mandatory active [ROLECE certificate](https://openmercantil.es/glosario/rolece) |
| **Harmonized EU Tender (`SARA`)** | Arts. 20–23 LCSP | `>= 143.000 €` (AGE) / `>= 221.000 €` / `>= 5.538.000 €` | Dual publication in PLACSP + EU TED (`DOUE`) |
| **Commercial Payment Cap** | Art. 4 Ley 3/2004 | **60 calendar days** maximum (`30d` default) | ECB reference rate + 8 pp late-payment interest + 40 € fee |
| **National Grants Disclosure** | Arts. 18 & 20 Ley 38/2003 | Corporate grants indexed by `CIF` | Published by IGAE in [BDNS](https://openmercantil.es/fuentes/bdns) |

## License

Code released under the **MIT License** (`LICENSE`). Public mercantile analytics and aggregate census figures derived from official Spanish state bulletins under **Ley 37/2007** and **CC BY 4.0** — attribution: [OpenMercantil (`openmercantil.es`)](https://openmercantil.es/).
