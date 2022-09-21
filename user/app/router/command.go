package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func Commands(rootCmd *cobra.Command, engine *gin.Engine) {
	// 	添加一个名为 job2 的命令，执行方式 go run main.go job2
	//var job2Cmd = &cobra.Command{
	//	Use:   "job2",
	//	Short: "This is a job to do yyy",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		run(engine, c.DemoJob2, args...)
	//	},
	//}
	//rootCmd.AddCommand(job2Cmd)
	//rootCmd.AddCommand(job2Cmd)
}
