package main

import ("fmt")

func main(){
	emojis := map[string]string{
		"red_angry_face": "\\uE416",
		"middle_finger": "\\U0001F595",
    	"unamused_face": "\\uE40E",
    	"slightly_smiling_face": "\\U0001F642",
    	"unamused_and_slightly_angry": "\\uE402",
    	"folded_hands": "\\uE41D",
    	"face_with_hand_over_mouth": "\\U0001F92D",
	}

	fmt.Println(emojis["red_angry_face"])
}