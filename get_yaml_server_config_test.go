package get_yaml_server_config

import "testing"

func TestGetConfig(t *testing.T) {
	expectedConfig := ServerConfig{
		Host:    "127.0.0.1",
		Port:    1235,
		Timeout: 1,
	}

	gotConfig, err := getConfig()
	if err != nil {
		t.Errorf("getConfig returned error")
	}

	if expectedConfig.Host != gotConfig.Host {
		t.Errorf("Wrong host. Expected: %s. Got: %s", expectedConfig.Host, gotConfig.Host)
	}
	if expectedConfig.Port != gotConfig.Port {
		t.Errorf("Wrong port. Expected: %d. Got: %d", expectedConfig.Port, gotConfig.Port)
	}
	if expectedConfig.Timeout != gotConfig.Timeout {
		t.Errorf("Wrong host. Expected: %f. Got: %f", expectedConfig.Timeout, gotConfig.Timeout)
	}
}
