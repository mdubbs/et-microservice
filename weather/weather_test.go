package weather

import (
  "testing"
)

func TestGetWeatherRecord(t *testing.T) {
  cases := []struct {
    in, want string
  }{
    {"48104", "Ann Arbor"},
    {"48864", "Okemos"},
    {"48823", "East Lansing"},
    {"48188", "Canton"},
  }
  for _, c := range cases {
    got, err := GetWeatherRecord(c.in)
    if err != nil {
      t.Errorf("GetWeatherRecord failed retriving weather for %q", c.in)
    }
    if got.Name != c.want {
      t.Errorf("GetWeatherRecord(%q) == %q, want %q", c.in, got.Name, c.want)
    }
  }
}