package main

import (
	"fmt"
	"math/rand"
	"time"
)

type AutomationScript struct {
	ID        string
	Name       string
	Triggers   []Trigger
	Actions    []Action
	Conditions []Condition
}

type Trigger struct {
	Type        string // e.g. " timer", " sensor", "input"
	Configuration map[string]interface{}
}

type Action struct {
	Type        string // e.g. "output", "api call", "notification"
	Configuration map[string]interface{}
}

type Condition struct {
	Type        string // e.g. "value comparison", "logic gate"
	Configuration map[string]interface{}
}

type Simulator struct {
	scripts    []AutomationScript
	sensors    []Sensor
	actors     []Actor
	triggers   map[string][]Trigger
	conditions map[string][]Condition
	actions     map[string][]Action
}

type Sensor struct {
	ID    string
	Value interface{}
}

type Actor struct {
	ID    string
	State interface{}
}

func (s *Simulator) Run() {
	for {
		// simulate sensor readings
		for _, sensor := range s.sensors {
			sensor.Value = rand.Intn(100)
		}

		// evaluate triggers
		for triggerID, triggers := range s.triggers {
			for _, trigger := range triggers {
				if trigger.Type == "timer" {
					trigger.Configuration["last-fired"] = time.Now().Unix()
				} else if trigger.Type == "sensor" {
					sensorID := trigger.Configuration["sensor-id"].(string)
					sensorValue := s.sensors[sensorID].Value
					if sensorValue.(int) > 50 {
						fmt.Printf("Trigger %s fired!\n", triggerID)
					}
				}
			}
		}

		// evaluate conditions
		for conditionID, conditions := range s.conditions {
			for _, condition := range conditions {
				if condition.Type == "value comparison" {
					sensorID := condition.Configuration["sensor-id"].(string)
					sensorValue := s.sensors[sensorID].Value
					if sensorValue.(int) > 50 {
						fmt.Printf("Condition %s is true!\n", conditionID)
					}
				}
			}
		}

		// execute actions
		for actionID, actions := range s.actions {
			for _, action := range actions {
				if action.Type == "output" {
					fmt.Printf("Action %s executed! Output: %v\n", actionID, action.Configuration["output"])
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	s := &Simulator{
		scripts:    []AutomationScript{},
		sensors:    []Sensor{},
		actors:     []Actor{},
		triggers:   map[string][]Trigger{},
		conditions: map[string][]Condition{},
		actions:     map[string][]Action{},
	}

	s.sensors = append(s.sensors, Sensor{ID: "temperature", Value: 0})
	s.sensors = append(s.sensors, Sensor{ID: "humidity", Value: 0})

	s.triggers["timer-1"] = []Trigger{
		{Type: "timer", Configuration: map[string]interface{}{"interval": 5}},
	}

	s.conditions["cond-1"] = []Condition{
		{Type: "value comparison", Configuration: map[string]interface{}{"sensor-id": "temperature", "operator": ">=", "value": 50}},
	}

	s.actions["action-1"] = []Action{
		{Type: "output", Configuration: map[string]interface{}{"output": "Temperature is high!"}},
	}

	s.Run()
}