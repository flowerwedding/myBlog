/**
 * @Title  output
 * @Description  日志格式化和输出
 * @Author  沈来
 * @Update  2020/8/3 21:50
 **/
package logger

import (
	"encoding/json"
	"time"
)

func (l *Logger) JSONFormat(message string) map[string]interface{}{
	data := make(Fields, len(l.fields) + 4)
	data["level"] = l.level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields{
			if _, ok := data[k]; !ok{
				data[k] = v
			}
		}
	}

	return data
}

func (l *Logger) Output(message string){
	body, _ := json.Marshal(l.JSONFormat(message))
	content := string(body)

	switch l.level{
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Print(content)
	case LevelPanic:
		l.newLogger.Print(content)
	}
}