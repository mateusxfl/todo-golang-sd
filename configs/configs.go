// Responsável por ler um arquivo de configuração, assim não precisamos setar configurações diretamente no código, evitando gerar binários novamente.

package configs

import (
	"github.com/spf13/viper"
)

// Ponteiro para a struct config, declarada a seguir.
var cfg *config

// Struct de configuração principal, que engloba duas structs secundárias, com configurações do servidor e do banco.
type config struct {
	API APIConfig
	DB  DBConfig
}

// Struct secundária, que engloba configurações do servidor.
type APIConfig struct {
	Port string
}

// Struct secundária, que engloba configurações do banco.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// A func init sempre é chamada no start das aplicações.
func init() {
	/*
		O viper é um package que auxilia na leitura do arquivo e setagem as configurações. Abaixo definiremos valores default para
		as nossas configurações.
	*/
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

/*
A func load vai ler o arquivo de configurações, assim, atribuindo valores as nossas structs. Caso não tenhamos nada, usaremos
os valores default da função anterior.
*/
func Load() error {
	// Nome do arquivo que iremos procurar.
	viper.SetConfigName("config")

	// Tipo de arquivo que iremos procurar.
	viper.SetConfigType("toml")

	// Onde esta o arquivo que iremos procurar.
	viper.AddConfigPath(".")

	// Ler o arquivo de configuração.
	err := viper.ReadInConfig()

	/*
		Caso tenhamos um erro, e este erro seja do tipo ConfigFileNotFoundError, poderemos continuar, pois se apenas não tivermos encontrado
		o arquivo estara tudo bem, pois ainda temos os valores default setados previamente em init.
	*/
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// Cria um ponteiro da nossa struct.
	cfg = new(config)

	// Atribui valores a struct referente a configuração do servidor.
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	// Atribui valores a struct referente a configuração do banco de dados.
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

// Retorna as configurações do banco.
func GetDB() DBConfig {
	return cfg.DB
}

// Retorna as configurações do servidor.
func GetServerPort() string {
	return cfg.API.Port
}
