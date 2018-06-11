package bot

import (
	"log"
	"strconv"
	"strings"

	"github.com/TinyKitten/sugoibot/dao"

	"github.com/TinyKitten/sugoibot/extapi"
	"github.com/nlopes/slack"
)

func (b *Bot) handleTodo(ev *slack.MessageEvent) error {
	op := strings.Replace(ev.Text, "./todo ", "", 1)
	opSpaces := strings.Fields(op)
	if len(opSpaces) < 1 {
		noOpsMsg := "オペレーションを指定してください。"
		b.handleCmdError(ev, noOpsMsg)
		return nil
	}

	member, err := extapi.GetMemberBySlackID(ev.User)
	if err != nil {
		return err
	}

	if strings.Index(op, "add") != -1 {
		taskName := strings.Replace(op, "add ", "", 1)
		taskNameSpaces := strings.Fields(taskName)
		if len(taskNameSpaces) != 1 {
			noTaskNameMsg := "タスク名を指定してください。"
			b.handleCmdError(ev, noTaskNameMsg)
			return nil
		}

		tid, err := dao.AddTask(member.Code, taskName)
		if err != nil {
			b.handleCmdError(ev, "整数値で指定してください。")
			return nil
		}

		dummyMsg := strconv.FormatInt(tid, 10) + "\n" + taskName + "\n" + member.Name + "(" + member.Code + ")"
		b.handleCmdCompletedWithPretext(ev, dummyMsg, "タスクを追加しました。")
		return nil
	}

	if strings.Index(op, "del") != -1 {
		taskID := strings.Replace(op, "del ", "", 1)
		taskIDSpaces := strings.Fields(taskID)
		if len(taskIDSpaces) != 1 {
			noIdmsg := "タスクIDを指定してください。"
			b.handleCmdError(ev, noIdmsg)
			return nil
		}

		tid, err := strconv.ParseInt(taskID, 10, 64)
		if err != nil {
			b.handleCmdError(ev, "整数値で指定してください。")
			return nil
		}

		task, err := dao.GetTaskByID(tid)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		if task == nil {
			b.handleCmdError(ev, "や、そんなタスクのNASA✋")
			return nil
		}

		err = dao.DeteleTask(tid)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		dummyMsg := taskID + "\n" + task.TaskName + "\n" + member.Name + "(" + member.Code + ")"
		b.handleCmdErrorWithPretext(ev, dummyMsg, "タスクを削除しました。")
		return nil
	}

	if strings.Index(op, "undone") != -1 {
		taskIDStr := strings.Replace(op, "undone ", "", 1)
		taskIDSpaces := strings.Fields(taskIDStr)
		if len(taskIDSpaces) != 1 {
			noIdmsg := "タスクIDを指定してください。"
			b.handleCmdError(ev, noIdmsg)
			return nil
		}

		tid, err := strconv.ParseInt(taskIDStr, 10, 64)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		task, err := dao.GetTaskByID(tid)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		if task == nil {
			b.handleCmdError(ev, "や、そんなタスクのNASA✋")
			return nil
		}

		err = dao.UpdateTaskStatus(tid, false)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		dummyMsg := taskIDStr + ": " + task.TaskName + "\n" + member.Name + "(" + member.Code + ")"
		b.handleCmdCompletedWithPretext(ev, dummyMsg, "タスクを未完了に設定しました。")
		return nil
	}

	if strings.Index(op, "done") != -1 {
		taskIDStr := strings.Replace(op, "done ", "", 1)
		taskIDSpaces := strings.Fields(taskIDStr)
		if len(taskIDSpaces) != 1 {
			noIdmsg := "タスクIDを指定してください。"
			b.handleCmdError(ev, noIdmsg)
			return nil
		}

		tid, err := strconv.ParseInt(taskIDStr, 10, 64)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		task, err := dao.GetTaskByID(tid)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		if task == nil {
			b.handleCmdError(ev, "や、そんなタスクのNASA✋")
			return nil
		}

		err = dao.UpdateTaskStatus(tid, true)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		dummyMsg := taskIDStr + ": " + task.TaskName + "\n" + member.Name + "(" + member.Code + ")"
		b.handleCmdCompletedWithPretext(ev, dummyMsg, "タスクを完了に設定しました。")
		return nil
	}

	if strings.Index(op, "mylist") != -1 {
		dummyPretext := member.Name + "(" + member.Code + ")のタスク一覧"
		todos, err := dao.GetTaskByMemberCode(member.Code)
		if err != nil {
			log.Println(err)
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		msg := ""

		for _, task := range *todos {
			member, err := extapi.GetMemberByCode(task.MemberCode)
			if err != nil {
				b.handleCmdError(ev, "内部エラーが発生しました。")
				return nil
			}
			completedStr := ""
			if task.Completed {
				completedStr = "完了"
			} else {
				completedStr = "未完了"
			}
			msg = msg + "*" + strconv.FormatInt(task.ID, 10) + ": " + task.TaskName + "(" + completedStr + ")* " + member.Name + "(" + member.Code + ")"
		}

		if msg == "" {
			msg = "NASA🚀"
		}

		b.handleCmdCompletedWithPretext(ev, msg, dummyPretext)
		return nil
	}

	if strings.Index(op, "all") != -1 {
		tasks, err := dao.GetAllTask()
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		dummyPretext := "メンバー全員のタスク"

		msg := ""

		for _, task := range *tasks {
			member, err := extapi.GetMemberByCode(task.MemberCode)
			if err != nil {
				b.handleCmdError(ev, "内部エラーが発生しました。")
				return nil
			}
			completedStr := ""
			if task.Completed {
				completedStr = "完了"
			} else {
				completedStr = "未完了"
			}

			msg = msg + "*" + strconv.FormatInt(task.ID, 10) + ": " + task.TaskName + "(" + completedStr + ")* " + member.Name + "(" + member.Code + ")"
		}

		if msg == "" {
			msg = "NASA🚀"
		}

		b.handleCmdCompletedWithPretext(ev, msg, dummyPretext)
		return nil
	}

	notImplMsg := "オペレーションが未定義です。"
	b.handleCmdError(ev, notImplMsg)
	return nil
}
