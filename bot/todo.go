package bot

import (
	"log"
	"strconv"

	"github.com/TinyKitten/sugoibot/constant"
	"github.com/TinyKitten/sugoibot/env"

	"github.com/TinyKitten/sugoibot/dao"

	"github.com/TinyKitten/sugoibot/extapi"
	"github.com/nlopes/slack"
)

func (b *Bot) handleTodo(ev *slack.MessageEvent, args ...string) error {
	if len(args) == 0 {
		noOpsMsg := "オペレーションを指定してください。"
		b.handleCmdError(ev, noOpsMsg)
		return nil
	}

	member, err := extapi.GetMemberBySlackID(ev.User)
	if err != nil {
		return err
	}

	adm, err := extapi.GetMemberBySlackID(env.GetAdminID())
	if err != nil {
		return err
	}

	isAdmin := member.Code == adm.Code

	op := args[0]

	switch op {
	case "add":
		if len(args) == 1 {
			noTaskNameMsg := "タスク名を指定してください。"
			b.handleCmdError(ev, noTaskNameMsg)
			return nil
		}

		tid, err := dao.AddTask(member.Code, args[1])
		if err != nil {
			b.handleCmdError(ev, "整数値で指定してください。")
			return nil
		}

		dummyMsg := strconv.FormatInt(tid, 10) + "\n" + args[1] + "\n" + member.Name + "(" + member.Code + ")"
		b.handleCmdCompletedWithPretext(ev, dummyMsg, "タスクを追加しました。")
		return nil
	case "del":
		if len(args) == 1 {
			noIdmsg := "タスクIDを指定してください。"
			b.handleCmdError(ev, noIdmsg)
			return nil
		}

		tid, err := strconv.ParseInt(args[1], 10, 64)
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

		if task.MemberCode != member.Code && !isAdmin {
			b.handleCmdError(ev, "や、貴様にそんな権限のNASA✋")
			return nil
		}

		taskAuthor, err := extapi.GetMemberByCode(task.MemberCode)
		switch {
		case err == constant.ERR_MEMBER_NOT_FOUND:
			taskAuthor.Name = "データベースから削除済み"
			taskAuthor.Code = task.MemberCode
		case err != nil:
			b.handleCmdError(ev, err.Error())
			return nil
		}

		err = dao.DeteleTask(tid)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		dummyMsg := args[1] + "\n" + task.TaskName + "\n" + taskAuthor.Name + "(" + taskAuthor.Code + ")"
		b.handleCmdErrorWithPretext(ev, dummyMsg, "タスクを削除しました。")
		return nil
	case "undone":
		if len(args) == 1 {
			noIdmsg := "タスクIDを指定してください。"
			b.handleCmdError(ev, noIdmsg)
			return nil
		}

		tid, err := strconv.ParseInt(args[1], 10, 64)
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

		if task.MemberCode != member.Code && !isAdmin {
			b.handleCmdError(ev, "や、貴様にそんな権限のNASA✋")
			return nil
		}

		taskAuthor, err := extapi.GetMemberByCode(task.MemberCode)
		switch {
		case err == constant.ERR_MEMBER_NOT_FOUND:
			taskAuthor.Name = "データベースから削除済み"
			taskAuthor.Code = task.MemberCode
		case err != nil:
			b.handleCmdError(ev, err.Error())
			return nil
		}

		err = dao.UpdateTaskStatus(tid, false)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		dummyMsg := args[1] + ": " + task.TaskName + "\n" + taskAuthor.Name + "(" + taskAuthor.Code + ")"
		b.handleCmdCompletedWithPretext(ev, dummyMsg, "タスクを未完了に設定しました。")
		return nil
	case "done":
		if len(args) == 1 {
			noIdmsg := "タスクIDを指定してください。"
			b.handleCmdError(ev, noIdmsg)
			return nil
		}

		tid, err := strconv.ParseInt(args[1], 10, 64)
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

		if task.MemberCode != member.Code && !isAdmin {
			b.handleCmdError(ev, "や、貴様にそんな権限のNASA✋")
			return nil
		}

		taskAuthor, err := extapi.GetMemberByCode(task.MemberCode)
		switch {
		case err == constant.ERR_MEMBER_NOT_FOUND:
			taskAuthor.Name = "データベースから削除済み"
			taskAuthor.Code = task.MemberCode
		case err != nil:
			b.handleCmdError(ev, err.Error())
			return nil
		}

		err = dao.UpdateTaskStatus(tid, true)
		if err != nil {
			b.handleCmdError(ev, "内部エラーが発生しました。")
			return nil
		}

		dummyMsg := args[1] + ": " + task.TaskName + "\n" + taskAuthor.Name + "(" + taskAuthor.Code + ")"
		b.handleCmdCompletedWithPretext(ev, dummyMsg, "タスクを完了に設定しました。")
		return nil
	case "mylist":
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
			msg = msg + "*" + strconv.FormatInt(task.ID, 10) + ": " + task.TaskName + "(" + completedStr + ")* " + member.Name + "(" + member.Code + ")\n"
		}

		if msg == "" {
			msg = "NASA🚀"
		}

		b.handleCmdCompletedWithPretext(ev, msg, dummyPretext)
		return nil
	case "all":
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

			msg = msg + "*" + strconv.FormatInt(task.ID, 10) + ": " + task.TaskName + "(" + completedStr + ")* " + member.Name + "(" + member.Code + ")\n"
		}

		if msg == "" {
			msg = "NASA🚀"
		}

		b.handleCmdCompletedWithPretext(ev, msg, dummyPretext)
		return nil
	default:
		notImplMsg := "オペレーションが未定義です。"
		b.handleCmdError(ev, notImplMsg)
		return nil
	}
}
