from telegram.ext import CommandHandler,MessageHandler, Filters,Updater
import logging
import requests
import time
import _thread
import os
import datetime
def updateMessage(update,chat_id,surl):
	today = datetime.date.today()
	while(True):
		time.sleep(1)
		try:
			r = requests.get(surl)
			for message in r.json():
				update.bot.send_message(chat_id=chat_id,text=message)
			r.close()
		except:
			pass
		if today != datetime.date.today():
			today = datetime.date.today()
			try:
				response = requests.post(os.environ.get("MCBOT_URL"), files={'input': (None,'save')})
				if response.status_code!=200:
					return
				if response.text!="":
					update.bot.send_document(chat_id=int(os.environ.get("MCBOT_CHAT_ID")),document=open(response.text,"rb"))
			except:
				pass
def save(update,context):
	try:
		response = requests.post(os.environ.get("MCBOT_URL"), files={'input': (None,'save 40M')})
		if response.status_code!=200:
			return
		if response.text!="":
			for filename in os.listdir(response.text):
				context.bot.send_document(chat_id=int(os.environ.get("MCBOT_CHAT_ID")),document=open(response.text+"/"+filename,"rb"))
	except:
		context.bot.send_message(chat_id=update.effective_chat.id, text="错误")

def info(update,context):
	context.bot.send_message(chat_id=update.effective_chat.id, text="chat_id : "+str(update.effective_chat.id)+"\nuser_id : "+str(update.message.from_user.id))
def start(update, context):
	context.bot.send_message(chat_id=update.effective_chat.id, text="I'm a bot, please talk to me!")

def ping(update,context):
	context.bot.send_message(chat_id=update.effective_chat.id, text="pong")

def sendMessage(update,context):
	try:
		response = requests.post(os.environ.get("MCBOT_URL"), files={'input': (None,'say '+
									update.message.from_user.first_name+" : "+
									update.message.text)})
		if response.status_code!=200:
			context.bot.send_message(chat_id=update.effective_chat.id, text="同步消息失败")
		if response.text!="":
				context.bot.send_message(chat_id=update.effective_chat.id, text=response.text)
	except:
		pass

def admin(update,context):
	if str(update.message.from_user.id) == os.environ.get("MCBOT_ADMIN_ID"):
		try:
			response = requests.post(os.environ.get("MCBOT_URL")+"/admin", files={'input': (None,' '.join(context.args))})
			if response.status_code!=200:
				context.bot.send_message(chat_id=update.effective_chat.id, text="发送命令失败")
			if response.text!="":
				context.bot.send_message(chat_id=update.effective_chat.id, text=response.text)
		except:
			context.bot.send_message(chat_id=update.effective_chat.id, text="发送命令失败")
	else:
		context.bot.send_message(chat_id=update.effective_chat.id, text="权限不够")
def unknown(update, context):
    context.bot.send_message(chat_id=update.effective_chat.id, text="Sorry, I didn't understand that command.")

def refistered(dispatcher):
	save_handler = CommandHandler('save',save)
	dispatcher.add_handler(save_handler)
	admin_handler = CommandHandler('admin',admin)
	dispatcher.add_handler(admin_handler)
	info_handler = CommandHandler('info',info)
	dispatcher.add_handler(info_handler)
	ping_handler = CommandHandler('ping',ping)
	dispatcher.add_handler(ping_handler)
	start_handler = CommandHandler('start',start)
	dispatcher.add_handler(start_handler)
	sendMessage_handler = MessageHandler(Filters.text & (~Filters.command), sendMessage)
	dispatcher.add_handler(sendMessage_handler)
	unknown_handler = MessageHandler(Filters.command, unknown)
	dispatcher.add_handler(unknown_handler)

def main(updater):
	_thread.start_new_thread(updateMessage,(updater,int(os.environ.get("MCBOT_CHAT_ID")),os.environ.get("MCBOT_URL")+"/message"))
	
	dispatcher = updater.dispatcher
	logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
					 level=logging.INFO)
	refistered(dispatcher)
	updater.start_polling()

if __name__=='__main__':
	try:
		updater = Updater(token=os.environ.get("MCBOT_TOKEN"),use_context=True)
		main(updater)
	except KeyboardInterrupt:
		try:
			updater.stop()
		except:
			pass