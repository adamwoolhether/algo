package algo

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPrintManager(t *testing.T) {

	input := []string{"First Document", "Second Document", "Third Document"}
	printer := newPrintMgr()

	for _, doc := range input {
		printer.queuePrintJob(doc)

		if top := printer.queue.peak(); top != input[0] {
			t.Errorf("exp top %s, got %s", input[0], top)
		}

		endOfLine := printer.queue.data[len(printer.queue.data)-1]
		if end := endOfLine; end != doc {
			t.Errorf("exp end of queue %s, got %s", endOfLine, end)
		}
	}

	var out bytes.Buffer

	printer.run(&out)

	scanner := bufio.NewScanner(&out)
	idx := 0
	for scanner.Scan() {
		if scanner.Text() != input[idx] {
			t.Errorf("Exp output %s, got %s", input[idx], scanner.Text())
		}
		idx++
	}

	queueLen := len(printer.queue.data)
	if queueLen != 0 {
		t.Errorf("Exp empyt print quue, got len %d", queueLen)
	}
}
