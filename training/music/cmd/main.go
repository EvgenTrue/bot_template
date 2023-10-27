//Задача 1: Музыкальная коллекция
//Предположим, у нас есть музыкальная коллекция, которая содержит различные типы медиа-файлов: альбомы, синглы и сборники. Коллекция предоставляет следующие возможности:
//*     Добавление медиа-файлов: Пользователь может добавлять новые медиа-файлы в коллекцию. Это могут быть альбомы, синглы или сборники.
//*     Просмотр списка медиа-файлов: Пользователь может просматривать список всех медиа-файлов в коллекции, включая информацию о названии, исполнителе и длительности.
//*     Расчет общей длительности: Коллекция может предоставлять информацию о общей длительности всех медиа-файлов, находящихся в ней.


//Пример взаимодействия:
//* Пользователь добавляет в коллекцию несколько альбомов различных исполнителей.
//* Пользователь просматривает список медиа-файлов в коллекции и видит информацию о названии альбомов и их исполнителях.
//* Пользователь просматривает общую длительность всех альбомов в коллекции.

package main
import "fmt"


type ItemTrack interface{
	Getname() string
	Getartist() string
	Getlen() int
	
}

type MusicCollection struct{
	Data []ItemTrack
	 
}

type Album struct{
	Name string
	Artist string
	time int
}
func (a *Album)Getname(){
	return a.Name
}

type Single struct{
	Name string
	Artist string
	time int
}
func (a *Album)Getname(){
	return a.Name
}
type Collection struct{
	Name string
	Artist string
	time int
}
func (c *CoCollection)Getname(){
	return s.Name
}

type Track struct{
	Name string
	Artist string
	time int
}	


func (m *MusicCollection)Addtrack(track ItemTrack){
 m.Album=append(m.Album,track)
}
func (m *MusicCollection)Addtrack(track ItemTrack){
	m.Album=append(m.Album,track)
   }
 
func (m *MusicCollection)ViewingNane(){

}

func (m *MusicCollection)VieewingTime(){

	
}


