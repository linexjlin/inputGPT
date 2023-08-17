package main

import (
	"fmt"

	"github.com/jeandeaual/go-locale"
)

var languages = map[string]map[string]string{
	"About": {
		"en": "About",
		"zh": "关于",
		"jp": "について",
		"es": "Acerca de",
		"de": "Über",
		"az": "حول",
	},
	"InputGPT": {
		"en": "InputGPT",
		"zh": "InputGPT",
		"jp": "InputGPT",
		"es": "InputGPT",
		"de": "InputGPT",
		"az": "InputGPT",
	},
	"InputGPT a Helpful input Assistant": {
		"en": "InputGPT",
		"zh": "InputGPT",
		"jp": "InputGPT",
		"es": "InputGPT, un asistente de entrada útil",
		"de": "InputGPT, ein hilfreicher Eingabeassistent",
		"az": "InputGPT, köməkçi daxil etmə asistenti",
	},
	"Exit": {
		"en": "Exit",
		"zh": "退出",
		"jp": "終了",
		"es": "Salir",
		"de": "Beenden",
		"az": "Çıxış",
	},
	"About the App": {

		"en": "About the App",
		"zh": "退出APP",
		"jp": "アプリについて",
		"es": "Acerca de la aplicación",
		"de": "Über die App",
		"az": "Tətbiq haqqında",
	},
	"Quit the whole app": {
		"zh": "退出APP",
		"jp": "アプリ全体を終了する",
		"es": "Salir de la aplicación",
		"de": "Beende die ganze App",
		"az": "Butun tətbiqi bağla",
	},
	"Click to active GPT": {
		"en": "Click to active GPT",
		"zh": "点击激活 GPT",
		"jp": "GPTをアクティブにするにはクリックしてください",
		"es": "Haz clic para activar GPT",
		"de": "Klicken Sie hier, um GPT zu aktivieren",
		"az": "GPT-ni aktivləşdirmək üçün klik edin",
	},
	"Clear Context": {
		"zh": "清除所有上下文",
		"jp": "すべてのコンテキストをクリアする",
		"es": "Borrar todo el contexto",
		"de": "Kontext löschen",
		"az": "Bütün konteksti təmizləyin",
	},
	"Default": {
		"zh": "默认",
		"jp": "デフォルト",
		"es": "Predeterminado",
		"de": "Standard",
		"az": "Varsayılan",
	},
	"Copy the question then click \"%s\" to query GPT": {
		"zh": "复制提问的文本，再按 \"%s\" 发送给GPT",
		"jp": "質問をコピーして、\"%s\"をクリックしてGPTにクエリを送信します",
		"es": "Copia la pregunta y luego haz clic en \"%s\" para consultar a GPT",
		"de": "Kopieren Sie die Frage und klicken Sie dann auf \"%s\", um GPT abzufragen",
		"az": "Sualı kopyalayın və sonra \"%s\" düyməsini vuraraq GPT-də sorğu göndərin",
	},
	"Clear Context %d/%d": {
		"zh": "清除所有上下文 %d/%d",
		"jp": "すべてのコンテキストをクリア %d/%d",
		"es": "Borrar todo el contexto %d/%d",
		"de": "Kontext löschen %d/%d",
		"az": "Bütün konteksti təmizlə %d/%d",
	},
	"Import": {
		"zh": "导入",
		"jp": "インポート",
		"es": "Importar",
		"de": "Importieren",
		"az": "İmport et",
	},
	"Manage Prompts": {
		"zh": "管理引导词",
		"jp": "プロンプトを管理する",
		"es": "Administrar los prompts",
		"de": "Prompts verwalten",
		"az": "Promptları idarə edin",
	},
	"Manager the prompts": {
		"zh": "打开文件夹，管理引导词",
		"jp": "プロンプトを管理する",
		"es": "Administrar los prompts",
		"de": "Prompts verwalten",
		"az": "Promptları idarə edin",
	},
	"Select this prompt": {
		"zh": "选择这个Prompt",
		"jp": "このプロンプトを選択",
		"es": "Seleccionar este prompt",
		"de": "Diesen Prompt auswählen",
		"az": "Bu promptu seçin",
	},
	"Open the project page": {
		"zh": "打开项目页面",
		"jp": "プロジェクトページを開く",
		"es": "Abrir la página del proyecto",
		"de": "Öffnen Sie die Projektseite",
		"az": "Proyekt səhifəsini açın",
	},
	"Set API KEY": {
		"zh": "设置 API KEY",
		"jp": "APIキーを設定する",
		"es": "Establecer la clave de API",
		"de": "API-Schlüssel festlegen",
		"az": "API AÇARI təyin edin",
	},
	"Modify, Delete prompts": {
		"zh": "更新、删除引导词",
		"jp": "プロンプトを変更、削除する",
		"es": "Modificar, eliminar los prompts",
		"de": "Prompts ändern, löschen",
		"az": "Promptları dəyişdirin, silin",
	},
	"Import a prompt from clipboard": {
		"zh": "把剪贴板上的Prompt导入",
		"jp": "クリップボードからプロンプトをインポートする",
		"es": "Importar un prompt desde el portapapeles",
		"de": "Importieren Sie einen Prompt aus der Zwischenablage",
		"az": "Promptu kliplikdən idxal edin",
	},
	"Set the OpenAI KEY, baseurl etc..": {
		"zh": "请设置OpenAI KEY，baseurl等参数。",
		"jp": "OpenAI KEY、baseurlなどを設定してください。",
		"es": "Establezca la clave de OpenAI, la URL base, etc.",
		"de": "Legen Sie den OpenAI KEY, die Base-URL usw. fest.",
		"az": "OpenAI AÇARI, baseurl və s. təyin edin.",
	},
}

var LANG = "az"

func setLang() {
	userLanguage, err := locale.GetLanguage()
	if err == nil {
		fmt.Println("Language:", userLanguage)
		LANG = userLanguage
	}
}

// give the right lanuage with query text when no match return the langText itself
func UText(langText string) string {
	if lang, ok := languages[langText]; ok {
		if text, ok := lang[LANG]; ok {
			return text
		}
	}
	if LANG != "en" {
		fmt.Printf("unmatch:\"%s\" \"%s\"\n", langText, LANG)
	}

	return langText
}
