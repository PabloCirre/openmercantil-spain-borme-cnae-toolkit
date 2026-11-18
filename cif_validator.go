// Package bormecnae implements Spanish corporate tax ID (CIF / NIF-jurídico)
// checksum verification under Royal Decree 1065/2007 (AEAT), EU VIES VAT formatting,
// and LCSP 9/2017 public procurement & Ley 3/2004 B2B trade credit thresholds.
//
// Reference documentation and live BORME / PLACSP / BDNS public registry lookup:
// https://openmercantil.es/como-comprobar-cif-empresa
package bormecnae

import (
	"fmt"
	"strings"
	"unicode"
)

// EntityForm describes the legal form associated with a Spanish CIF prefix letter.
type EntityForm struct {
	Letter      byte
	Description string
	ControlMode string // "digit", "letter", or "either"
}

var cifLetters = map[byte]EntityForm{
	'A': {'A', "Sociedad Anónima (S.A.)", "digit"},
	'B': {'B', "Sociedad de Responsabilidad Limitada (S.L. / S.L.U.)", "digit"},
	'C': {'C', "Sociedad Colectiva", "either"},
	'D': {'D', "Sociedad Comanditaria", "either"},
	'E': {'E', "Comunidad de Bienes / Herencia Yacente", "digit"},
	'F': {'F', "Sociedad Cooperativa", "either"},
	'G': {'G', "Asociación / Fundación", "either"},
	'H': {'H', "Comunidad de Propietarios en Propiedad Horizontal", "digit"},
	'J': {'J', "Sociedad Civil", "either"},
	'N': {'N', "Entidad Extranjera sin Establecimiento Permanente", "letter"},
	'P': {'P', "Corporación Local (Ayuntamiento / Diputación)", "letter"},
	'Q': {'Q', "Organismo Público / Entidad Estatal", "letter"},
	'R': {'R', "Congregación o Institución Religiosa", "letter"},
	'S': {'S', "Órgano de la Administración del Estado o CC.AA.", "letter"},
	'U': {'U', "Unión Temporal de Empresas (U.T.E.)", "digit"},
	'V': {'V', "Agrupación de Interés Económico / Fondo / Otros", "either"},
	'W': {'W', "Establecimiento Permanente de Entidad No Residente", "letter"},
}

const controlMap = "JABCDEFGHI"

// ValidateCIF validates a 9-character Spanish corporate CIF/NIF according to RD 1065/2007.
func ValidateCIF(raw string) (bool, EntityForm, error) {
	cif := strings.ToUpper(strings.TrimSpace(strings.ReplaceAll(raw, "-", "")))
	cif = strings.TrimPrefix(cif, "ES")
	if len(cif) != 9 {
		return false, EntityForm{}, fmt.Errorf("expected 9 characters, got %d", len(cif))
	}
	form, ok := cifLetters[cif[0]]
	if !ok {
		return false, EntityForm{}, fmt.Errorf("invalid CIF organization prefix %q", cif[0])
	}
	evenSum := 0
	oddSum := 0
	for i := 1; i <= 7; i++ {
		ch := rune(cif[i])
		if !unicode.IsDigit(ch) {
			return false, form, fmt.Errorf("position %d must be numeric", i)
		}
		d := int(cif[i] - '0')
		if i%2 == 0 {
			evenSum += d
		} else {
			dbl := d * 2
			oddSum += (dbl / 10) + (dbl % 10)
		}
	}
	total := evenSum + oddSum
	unit := (10 - (total % 10)) % 10
	expectedDigit := byte('0' + unit)
	expectedLetter := controlMap[unit]
	actual := cif[8]

	switch form.ControlMode {
	case "digit":
		return actual == expectedDigit, form, nil
	case "letter":
		return actual == expectedLetter, form, nil
	default:
		return actual == expectedDigit || actual == expectedLetter, form, nil
	}
}

// FormatEUVIES returns the intra-Community EU VAT identifier (ES + CIF) if valid.
func FormatEUVIES(raw string) (string, error) {
	ok, _, err := ValidateCIF(raw)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("invalid CIF checksum")
	}
	clean := strings.TrimPrefix(strings.ToUpper(strings.TrimSpace(strings.ReplaceAll(raw, "-", ""))), "ES")
	return "ES" + clean, nil
}
