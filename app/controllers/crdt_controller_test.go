package controllers

/*
func TestGetCrdtValue(t *testing.T) {
	s := NewServer("dummy", "libp2p", "dummy", false)
	rr := httptest.NewRecorder()

	// Do sub request
	topic := "randomTopic1"
	data, err := json.Marshal(SubRequestBody{
		Topic: topic,
	})
	if err != nil {
		t.Error(err)
	}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest(http.MethodPost, "", body)
	if err != nil {
		t.Error(err)
	}
	makeHTTPHandler(s.handleSubscribe)(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", rr.Result().StatusCode)
	}

	// Test if Topics include the topics after sub
	expected := fmt.Sprintf("\"Subscribed to the topic %s\"", topic)
	b, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Error(err)
	}
	if strings.TrimRight(string(b), "\n") != expected {
		t.Errorf("expected %s but got %s", expected, string(b))
	}
}
*/
