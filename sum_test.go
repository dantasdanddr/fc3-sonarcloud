package main

import (
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("teste saida", func(t *testing.T) {
		cmd := exec.Command("go", "run", "sum.go")
		output, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("Erro ao executar comando: %v", err)
		}

		expected := "Olá, mundo!\n"
		if string(output) != expected {
			t.Errorf("Saída inesperada: %v", string(output))
		}
	})
}

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}

func TestSub(t *testing.T) {

	result := sub(3, 1)

	if result != 2 {
		t.Error("The result must be 2")
	}
}

func TestTimes(t *testing.T) {

	result := times(3, 3)

	if result != 9 {
		t.Error("The result must be 9")
	}
}

func TestSumX(t *testing.T) {

	result := sumX(3, 2)

	if result != 8 {
		t.Error("The result must be 8")
	}
}
