package model

func GetPeopleByID(id string)(error,*People)  {
   return nil,&People{
	   ID:        "12134",
	   Name:      "test",
	   BirthYear: nil,
	   EyeColor:  nil,
	   Gender:    nil,
	   HairColor: nil,
	   Height:    nil,
	   Mass:      nil,
	   SkinColor: nil,
	   Films:     nil,
   }
}
