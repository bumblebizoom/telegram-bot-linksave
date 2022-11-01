package telegram

const msgHelp = `Я можу зберігати і вести ваші сторінки. Також я можу запропонувати вам їх прочитати.

Для того, щоб зберегти сторінку, просто надішліть мені всі посилання на неї!

Для того, щоб отримати випадкову сторінку з вашого списку, надішліть мені команду /cmd
Увага! Після цього, ця сторінка буде вилучена з вашого списку!!!`
const msgHello = "Привііііт! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Не розумію що робитиии 🥶"
	msgNoSavedPages   = "У тебе немає збережених сторінок 🙄"
	msgSaved          = "Зберіг! ✅"
	msgAlreadyExists  = "Ця сторінка вже є в списку 🤗"
)
