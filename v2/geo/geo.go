package geo

import "errors"

func CubeVolume(n int) (int,error) {
	if n!=0{
	return n * n * n,nil
}else{
	return 0, errors.New("Zero Value not valid")
}

}
