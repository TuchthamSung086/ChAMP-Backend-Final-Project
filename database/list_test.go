package database_test

import (
	"ChAMP-Backend-Final-Project/database"
	"ChAMP-Backend-Final-Project/initializers"
	"ChAMP-Backend-Final-Project/models"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test List Service, Test Task Service
var tls database.ListService
var tts database.TaskService

func init() {
	// Get environment variables
	initializers.LoadEnvVariables()
	// Initialize database connection
	db, err := initializers.ConnectToDB() // Note: I know I shouldn't use the real database to test, but I don't have the time anymore to do Docker stuff AHHH
	if err != nil {
		panic("Can't connect to database")
	}
	tls = database.NewListService(db)
	tts = database.NewTaskService(db)
}

// Get all function must return nothing for an empty database
func TestGetAllListEmptyDB(t *testing.T) {
	assert := assert.New(t)
	// Clear Database
	tts.DeleteAll()
	tls.DeleteAll()

	lists, err := tls.GetAll()

	assert.Nil(err) // Must not error
	assert.Equal(len(lists), 0)
}

// Create this List and Task and check
// L0: T0
// L1: T1 T2
// L2: T3 T4 T5
func TestCreate(t *testing.T) {
	assert := assert.New(t)
	// ClearDB
	tts.DeleteAll()
	tls.DeleteAll()

	// remember ids of lists and tasks
	var listIds [10]uint
	var taskIds [10]uint

	// temp variables
	var cl *models.ControllerList
	var ct *models.ControllerTask
	var err error

	// ***** Test CREATE *****
	// Create 3 lists --> those lists must exist and have orders as 1 2 3
	for i := 0; i <= 2; i++ {
		cl, err = tls.Create(&models.ControllerList{Title: strconv.Itoa(i)}) // List 0 1 2
		assert.Equal(cl.Order, i+1)                                          // Order 1 2 3
		assert.Nil(err)                                                      // No error
		listIds[i] = uint(cl.ID)                                             // Save Id
	}

	// Create Tasks like this
	// L0: T0
	// L1: T1 T2
	// L2: T3 T4 T5
	// Create 1 task in List 0 --> Must be in order 1
	ct, err = tts.Create(&models.ControllerTask{Title: "0", ListID: uint(listIds[0])}) // Task 0
	assert.Equal(ct.Order, 1)                                                          // Order 1
	assert.Nil(err)                                                                    // No error
	taskIds[0] = uint(ct.ID)                                                           // Save Id
	// Create 2 tasks in List 1 --> Must be in order 1 2
	for i := 1; i < 3; i++ {
		ct, err = tts.Create(&models.ControllerTask{Title: strconv.Itoa(i), ListID: uint(listIds[1])}) // Task 1 2
		assert.Equal(ct.Order, i)                                                                      // Order 1 2
		assert.Nil(err)                                                                                // No error
		taskIds[i] = uint(ct.ID)                                                                       // Save Id
	}
	// Create 3 tasks in List 2 --> Must be in order 1 2 3
	for i := 3; i < 6; i++ {
		ct, err = tts.Create(&models.ControllerTask{Title: strconv.Itoa(i), ListID: uint(listIds[2])}) // Task 3 4 5
		assert.Equal(ct.Order, i-2)                                                                    // Order 1 2 3
		assert.Nil(err)                                                                                // No error
		taskIds[i] = uint(ct.ID)                                                                       // Save Id
	}
	// ClearDB
	tts.DeleteAll()
	tls.DeleteAll()
}

// Create this List and Task and get them
// L0: T0
// L1: T1 T2
// L2: T3 T4 T5
func TestGet(t *testing.T) {
	assert := assert.New(t)
	// ClearDB
	tts.DeleteAll()
	tls.DeleteAll()

	// remember ids of lists and tasks
	var listIds [10]uint
	var taskIds [10]uint

	// temp variables
	var cl *models.ControllerList
	var ct *models.ControllerTask
	var err error

	// ***** CREATE *****
	// Create 3 lists --> those lists must exist and have orders as 1 2 3
	for i := 0; i <= 2; i++ {
		cl, err = tls.Create(&models.ControllerList{Title: strconv.Itoa(i)}) // List 0 1 2
		listIds[i] = uint(cl.ID)                                             // Save Id
	}

	// Create Tasks like this
	// L0: T0
	// L1: T1 T2
	// L2: T3 T4 T5
	// Create 1 task in List 0 --> Must be in order 1
	ct, err = tts.Create(&models.ControllerTask{Title: "0", ListID: uint(listIds[0])}) // Task 0
	taskIds[0] = uint(ct.ID)                                                           // Save Id
	// Create 2 tasks in List 1 --> Must be in order 1 2
	for i := 1; i < 3; i++ {
		ct, err = tts.Create(&models.ControllerTask{Title: strconv.Itoa(i), ListID: uint(listIds[1])}) // Task 1 2
		taskIds[i] = uint(ct.ID)                                                                       // Save Id
	}
	// Create 3 tasks in List 2 --> Must be in order 1 2 3
	for i := 3; i < 6; i++ {
		ct, err = tts.Create(&models.ControllerTask{Title: strconv.Itoa(i), ListID: uint(listIds[2])}) // Task 3 4 5
		taskIds[i] = uint(ct.ID)                                                                       // Save Id
	}

	// ***** Test Get *****
	lists, err := tls.GetAll()
	assert.Nil(err)
	for _, list := range lists {
		if list.ID == listIds[0] { // List 0
			listFromId, err := tls.GetById(listIds[0]) // Must be same as Get By Id
			assert.Nil(err)
			assert.Equal(list, listFromId)
			assert.EqualValues(len(list.Tasks), 1)                     // Have 1 Task
			assert.EqualValues(int(list.Tasks[0].ID), int(taskIds[0])) // Task Id == 0
			taskFromId, err := tts.GetById(list.Tasks[0].ID)
			assert.Nil(err)
			assert.Equal(list.Tasks[0], taskFromId) // Must be same as Get By Id
		} else if list.ID == listIds[1] { // List 1
			listFromId, err := tls.GetById(listIds[1]) // Must be same as Get By Id
			assert.Nil(err)
			assert.Equal(list, listFromId)
			assert.EqualValues(len(list.Tasks), 2)                                                                             // Have 2 Tasks
			assert.ElementsMatch([]int{int(taskIds[1]), int(taskIds[2])}, []int{int(list.Tasks[0].ID), int(list.Tasks[1].ID)}) // Task Id == 1 or 2
		} else if list.ID == listIds[2] { // List 2
			listFromId, err := tls.GetById(listIds[2]) // Must be same as Get By Id
			assert.Nil(err)
			assert.EqualValues(list, listFromId)
			assert.EqualValues(len(list.Tasks), 3)                                                                                                                     // Have 3 Tasks
			assert.ElementsMatch([]int{int(taskIds[3]), int(taskIds[4]), int(taskIds[5])}, []int{int(list.Tasks[0].ID), int(list.Tasks[1].ID), int(list.Tasks[2].ID)}) // Task Id == 1 or 2
		}
		var tasks []*models.ControllerTask
		tasks, err = tts.GetAll()
		assert.Nil(err)
		var m = map[int]int{0: 0, 1: 1, 2: 1, 3: 2, 4: 2, 5: 2} // Map Task number to List number
		var taskNum int
		for _, task := range tasks {
			taskById, err := tts.GetById(task.ID) // Must be same with get by id
			assert.Nil(err)
			assert.EqualValues(task, taskById)
			taskNum, err = strconv.Atoi(task.Title)
			assert.Nil(err)
			assert.EqualValues(int(task.ListID), int(listIds[m[taskNum]]))
		}
	}
	// ClearDB
	tts.DeleteAll()
	tls.DeleteAll()
}

// Move T0 To L2 TaskOrder 2, Move L1 to ListOrder 1
// FROM
// L0: T0
// L1: T1 T2
// L2: T3 T4 T5
// TO
// L1: T1 T2
// L0:
// L2: T3 T0 T4 T5
func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	// ClearDB
	tts.DeleteAll()
	tls.DeleteAll()

	// remember ids of lists and tasks
	var listIds [10]uint
	var taskIds [10]uint

	// temp variables
	var cl *models.ControllerList
	var ct *models.ControllerTask
	var err error

	// ***** CREATE *****
	// Create 3 lists --> those lists must exist and have orders as 1 2 3
	for i := 0; i <= 2; i++ {
		cl, err = tls.Create(&models.ControllerList{Title: strconv.Itoa(i)}) // List 0 1 2
		listIds[i] = uint(cl.ID)                                             // Save Id
	}

	// Create Tasks like this
	// L0: T0
	// L1: T1 T2
	// L2: T3 T4 T5
	// Create 1 task in List 0 --> Must be in order 1
	ct, err = tts.Create(&models.ControllerTask{Title: "0", ListID: uint(listIds[0])}) // Task 0
	taskIds[0] = uint(ct.ID)                                                           // Save Id
	// Create 2 tasks in List 1 --> Must be in order 1 2
	for i := 1; i < 3; i++ {
		ct, err = tts.Create(&models.ControllerTask{Title: strconv.Itoa(i), ListID: uint(listIds[1])}) // Task 1 2
		taskIds[i] = uint(ct.ID)                                                                       // Save Id
	}
	// Create 3 tasks in List 2 --> Must be in order 1 2 3
	for i := 3; i < 6; i++ {
		ct, err = tts.Create(&models.ControllerTask{Title: strconv.Itoa(i), ListID: uint(listIds[2])}) // Task 3 4 5
		taskIds[i] = uint(ct.ID)                                                                       // Save Id
	}

	// ***** TEST UPDATE *****
	// Move T0 To L2 TaskOrder 2
	// L0:
	// L1: T1 T2
	// L2: T3 T0 T4 T5
	tts.Update(taskIds[0], &models.ControllerTask{ListID: listIds[2], Order: 2})
	// List0 must be empty
	list0, err := tls.GetById(listIds[0])
	assert.EqualValues(len(list0.Tasks), 0)
	var orders = map[int]int{0: 2, 3: 1, 4: 3, 5: 4} // New Order of Tasks in List 2
	tasks, err := tts.GetAll()
	assert.Nil(err)
	for _, task := range tasks {
		if task.ListID == listIds[2] {
			assert.Equal(int(listIds[2]), int(task.ListID)) // Must be in List 2
			taskNum, err := strconv.Atoi(task.Title)
			assert.Nil(err)
			assert.Equal(task.Order, orders[taskNum]) // Must be in right order
		}
	}

	// Move L1 to ListOrder 1
	// L1: T1 T2
	// L0:
	// L2: T3 T0 T4 T5
	tls.Update(listIds[1], &models.ControllerList{Order: 1})
	var listOrders = map[int]int{0: 2, 1: 1, 2: 3} // New Order of Lists
	lists, err := tls.GetAll()
	assert.Nil(err)
	for _, list := range lists {
		listNum, _ := strconv.Atoi(list.Title)

		assert.EqualValues(list.Order, listOrders[listNum])
	}

	// ClearDB
	tts.DeleteAll()
	tls.DeleteAll()
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	// ClearDB
	tts.DeleteAll()
	tls.DeleteAll()

	// remember ids of lists and tasks
	var listIds [10]uint
	var taskIds [10]uint

	// temp variables
	var cl *models.ControllerList
	var ct *models.ControllerTask
	var err error

	// ***** CREATE *****
	// Create 3 lists --> those lists must exist and have orders as 1 2 3
	for i := 0; i <= 2; i++ {
		cl, err = tls.Create(&models.ControllerList{Title: strconv.Itoa(i)}) // List 0 1 2
		listIds[i] = uint(cl.ID)                                             // Save Id
	}

	// Create Tasks like this
	// L0: T0
	// L1: T1 T2
	// L2: T3 T4 T5
	// Create 1 task in List 0 --> Must be in order 1
	ct, err = tts.Create(&models.ControllerTask{Title: "0", ListID: uint(listIds[0])}) // Task 0
	taskIds[0] = uint(ct.ID)                                                           // Save Id
	// Create 2 tasks in List 1 --> Must be in order 1 2
	for i := 1; i < 3; i++ {
		ct, err = tts.Create(&models.ControllerTask{Title: strconv.Itoa(i), ListID: uint(listIds[1])}) // Task 1 2
		taskIds[i] = uint(ct.ID)                                                                       // Save Id
	}
	// Create 3 tasks in List 2 --> Must be in order 1 2 3
	for i := 3; i < 6; i++ {
		ct, err = tts.Create(&models.ControllerTask{Title: strconv.Itoa(i), ListID: uint(listIds[2])}) // Task 3 4 5
		taskIds[i] = uint(ct.ID)                                                                       // Save Id
	}

	// *** TEST DELETE ***

	// Delete Task 3
	// L0: T0
	// L1: T1 T2
	// L2: T4 T5
	_, err = tts.Delete(taskIds[3])
	assert.Nil(err)
	t4, err := tts.GetById(taskIds[4]) // T4 Order 1
	assert.Nil(err)
	assert.EqualValues(t4.Order, 1)
	t5, err := tts.GetById(taskIds[5]) // T5 Order 2
	assert.Nil(err)
	assert.EqualValues(t5.Order, 2)

	// Delete L1
	_, err = tls.Delete(listIds[1])
	assert.Nil(err)
	_, err = tts.GetById(taskIds[1]) // T1 must not exist
	assert.NotNil(err)
	_, err = tts.GetById(taskIds[2]) // T2 must not exist
	assert.NotNil(err)
	l2, err := tls.GetById(listIds[2])
	assert.Nil(err)
	assert.EqualValues(l2.Order, 2) // L2 Order becomes 2 from 3
	l0, err := tls.GetById(listIds[0])
	assert.Nil(err)
	assert.EqualValues(l0.Order, 1) // L0 Order remains unchanged (1)

	// ClearDB
	tts.DeleteAll()
	tls.DeleteAll()
}
