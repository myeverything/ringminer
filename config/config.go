/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package config

import (
	"github.com/naoina/toml"
	"os"
	"reflect"
	"go.uber.org/zap"
)

func LoadConfig(file string) (*GlobalConfig, error) {
	if ("" == file) {
		dir, _ := os.Getwd()
		file = dir + "/config/ringminer.toml"
	}

	io, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer io.Close()

	c := &GlobalConfig{}
	c.defaultConfig()
	if err := toml.NewDecoder(io).Decode(c); err != nil {
		panic(err)
	}

	if _,err := validator(c);nil != err {
		return nil,err
	}
	return c, nil
}

type GlobalConfig struct {
	Title       string `required:"true"`
	Owner       struct {
		Name string
	}
	Database    DbOptions
	Ipfs        IpfsOptions
	ChainClient ChainClientOptions
	Miner       MinerOptions
	LogOptions  zap.Config
}

//todo:optimize it
func  validator(c *GlobalConfig) (bool,error) {
	//cv := reflect.ValueOf(c).Elem()
	//for i:=0;i<cv.NumField();i++ {
	//	println("field:",cv.Field(i).String(), " tag:",cv.Type().Field(i).Tag.Get("required"), " f:", isNil(cv.Field(i)))
	//}
	//if "true" == cv.Type().Field(0).Tag.Get("required") {
	//	if cv.Field(0).IsNil() {
	//		return false
	//	}
	//}
	return true,nil
}

func isNil(v reflect.Value) bool {
	switch v.Type().Kind() {
	case reflect.Invalid:
		return false
	case reflect.String:
		return v.String() != ""
	}
	return false;
}

func (c *GlobalConfig) defaultConfig() {

}

type IpfsOptions struct {
	Server string
	Port int
	Topic string
}

type DbOptions struct {
	Server string `required:"true"`
	Port int `required:"true"`
	Name string `required:"true"`
	DataDir string
	CacheCapacity int
	BufferCapacity int
}

type ChainClientOptions struct {
	RawUrl string	`required:"true"`
	Eth    struct{
		    GasPrice int
		    GasLimit int
		    PrivateKeys map[string]string `required:"true"` //地址 -> 加密后的私钥，如果密码不对，地址与私钥则不会匹配
		    Password string	//密码，用于加密私钥，最好不出现在配置文件中
	    }
}

type MinerOptions struct {
	LoopringImps []ContractOpts `required:"true"`
	LoopringFingerprints []ContractOpts `required:"true"`
	RingMaxLength	int
}

type ContractOpts struct {
	Abi	string `required:"true"`
	Address	string `required:"true"`
}
