// Code generated by oto; DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/pacedotdev/pace"
)

func runCommand(ctx context.Context, args []string) error {
	if len(args) < 2 {
		showHelp(args)
		return nil
	}
	switch args[1] {
	case "help":
		showHelp(args)
		return nil
	case "list":
		printList()
		return nil
	case "templates":
		printTemplates()
		return nil

	case "CardsService.CreateCard":
		return CardsServiceCreateCard(ctx, args)

	case "CardsService.GetCard":
		return CardsServiceGetCard(ctx, args)

	case "CommentsService.AddComment":
		return CommentsServiceAddComment(ctx, args)

	default:
		fmt.Println("unknown command:", args[1])
		showHelp(nil)
	}
	return nil
}

func showHelp(args []string) {
	if len(args) < 3 {
		printUsage()

		fmt.Print("* CardsService")
		commentForCardsService := `CardsService is used to work with cards.`
		if commentForCardsService != "" {
			fmt.Print(" - ", commentForCardsService)
		}
		fmt.Println()

		fmt.Print("* CommentsService")
		commentForCommentsService := ``
		if commentForCommentsService != "" {
			fmt.Print(" - ", commentForCommentsService)
		}
		fmt.Println()

		fmt.Println()
		fmt.Println("  pace help <service>[.<method>] - print specific help")
		fmt.Println("  pace list - list all services and methods")
		fmt.Println("  pace templates - show copy and paste examples")
		printFlagDefaults(args)
		return
	}
	showHelpFor(args, args[2])
}

func showHelpFor(args []string, service string) {
	switch service {

	case "CardsService":
		fmt.Printf("methods for %s:\n", service)

		fmt.Print("* CardsService.CreateCard")
		commentForCardsServiceCreateCard := `CreateCard creates a new Card.`
		if commentForCardsServiceCreateCard != "" {
			fmt.Print(" - ", commentForCardsServiceCreateCard)
		}
		fmt.Println()

		fmt.Print("* CardsService.GetCard")
		commentForCardsServiceGetCard := `GetCard gets a card.`
		if commentForCardsServiceGetCard != "" {
			fmt.Print(" - ", commentForCardsServiceGetCard)
		}
		fmt.Println()

	case "CardsService.CreateCard":
		fmt.Println(`Usage for CardsService.CreateCard

  pace CardsService.CreateCard [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		var (
			apikey string
			secret string
			host   string
			debug  bool
		)
		flags.StringVar(&apikey, "apikey", "", "Pace API Key")
		flags.StringVar(&secret, "secret", "", "Pace API secret")
		flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
		flags.BoolVar(&debug, "debug", false, "prints debug information")
		var request pace.CreateCardRequest
		addFlagsForCreateCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CardsService.GetCard":
		fmt.Println(`Usage for CardsService.GetCard

  pace CardsService.GetCard [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		var (
			apikey string
			secret string
			host   string
			debug  bool
		)
		flags.StringVar(&apikey, "apikey", "", "Pace API Key")
		flags.StringVar(&secret, "secret", "", "Pace API secret")
		flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
		flags.BoolVar(&debug, "debug", false, "prints debug information")
		var request pace.GetCardRequest
		addFlagsForGetCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CommentsService":
		fmt.Printf("methods for %s:\n", service)

		fmt.Print("* CommentsService.AddComment")
		commentForCommentsServiceAddComment := `AddComment adds a comment.`
		if commentForCommentsServiceAddComment != "" {
			fmt.Print(" - ", commentForCommentsServiceAddComment)
		}
		fmt.Println()

	case "CommentsService.AddComment":
		fmt.Println(`Usage for CommentsService.AddComment

  pace CommentsService.AddComment [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		var (
			apikey string
			secret string
			host   string
			debug  bool
		)
		flags.StringVar(&apikey, "apikey", "", "Pace API Key")
		flags.StringVar(&secret, "secret", "", "Pace API secret")
		flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
		flags.BoolVar(&debug, "debug", false, "prints debug information")
		var request pace.AddCommentRequest
		addFlagsForAddCommentRequest(flags, "", &request)
		flags.PrintDefaults()

	default:
		fmt.Println(service, "is not supported")
		showHelp(nil)
	}
}

func printUsage() {
	fmt.Println(`Usage:
  pace <service>.<method> [args...]`)
	fmt.Println()
}

func CardsServiceCreateCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		apikey string
		secret string
		host   string
		debug  bool
	)
	flags.StringVar(&apikey, "apikey", "", "Pace API Key")
	flags.StringVar(&secret, "secret", "", "Pace API secret")
	flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
	flags.BoolVar(&debug, "debug", false, "prints debug information")
	var request pace.CreateCardRequest
	addFlagsForCreateCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if apikey == "" {
		apikey = os.Getenv("PACE_API_KEY")
	}
	if apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if secret == "" {
		secret = os.Getenv("PACE_API_SECRET")
	}
	if secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(apikey, secret)
	client.RemoteHost = host
	if debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCardsService(client)
	resp, err := service.CreateCard(ctx, request)
	if err != nil {
		return err
	}
	printCreateCardResponse(resp)
	return nil
}

func CardsServiceGetCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		apikey string
		secret string
		host   string
		debug  bool
	)
	flags.StringVar(&apikey, "apikey", "", "Pace API Key")
	flags.StringVar(&secret, "secret", "", "Pace API secret")
	flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
	flags.BoolVar(&debug, "debug", false, "prints debug information")
	var request pace.GetCardRequest
	addFlagsForGetCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if apikey == "" {
		apikey = os.Getenv("PACE_API_KEY")
	}
	if apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if secret == "" {
		secret = os.Getenv("PACE_API_SECRET")
	}
	if secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(apikey, secret)
	client.RemoteHost = host
	if debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCardsService(client)
	resp, err := service.GetCard(ctx, request)
	if err != nil {
		return err
	}
	printGetCardResponse(resp)
	return nil
}

func CommentsServiceAddComment(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		apikey string
		secret string
		host   string
		debug  bool
	)
	flags.StringVar(&apikey, "apikey", "", "Pace API Key")
	flags.StringVar(&secret, "secret", "", "Pace API secret")
	flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
	flags.BoolVar(&debug, "debug", false, "prints debug information")
	var request pace.AddCommentRequest
	addFlagsForAddCommentRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if apikey == "" {
		apikey = os.Getenv("PACE_API_KEY")
	}
	if apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if secret == "" {
		secret = os.Getenv("PACE_API_SECRET")
	}
	if secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(apikey, secret)
	client.RemoteHost = host
	if debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCommentsService(client)
	resp, err := service.AddComment(ctx, request)
	if err != nil {
		return err
	}
	printAddCommentResponse(resp)
	return nil
}

func printFlagDefaults(args []string) {
	fmt.Println()
	fmt.Println("Flags:")
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		apikey string
		host   string
		debug  bool
	)
	flags.StringVar(&apikey, "apikey", "", "Pace API Key")
	flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
	flags.BoolVar(&debug, "debug", false, "prints debug information")
	flags.PrintDefaults()
}

func printList() {

	fmt.Printf("CardsService.CreateCard")
	commentCardsServiceCreateCard := `CreateCard creates a new Card.`
	if len(commentCardsServiceCreateCard) > 0 {
		fmt.Printf(" - %s", commentCardsServiceCreateCard)
	}
	fmt.Println()

	fmt.Printf("CardsService.GetCard")
	commentCardsServiceGetCard := `GetCard gets a card.`
	if len(commentCardsServiceGetCard) > 0 {
		fmt.Printf(" - %s", commentCardsServiceGetCard)
	}
	fmt.Println()

	fmt.Printf("CommentsService.AddComment")
	commentCommentsServiceAddComment := `AddComment adds a comment.`
	if len(commentCommentsServiceAddComment) > 0 {
		fmt.Printf(" - %s", commentCommentsServiceAddComment)
	}
	fmt.Println()

}

func printTemplates() {

	fmt.Printf("pace CardsService.CreateCard ")
	printArgslistCreateCardRequest()
	fmt.Println()

	fmt.Printf("pace CardsService.GetCard ")
	printArgslistGetCardRequest()
	fmt.Println()

	fmt.Printf("pace CommentsService.AddComment ")
	printArgslistAddCommentRequest()
	fmt.Println()

}

func addFlagsForAddCommentRequest(flags *flag.FlagSet, prefix string, v *pace.AddCommentRequest) {

	flags.StringVar(&v.OrgID, prefix+"OrgID", "", ``)

	flags.StringVar(&v.TargetKind, prefix+"TargetKind", "", ``)

	flags.StringVar(&v.TargetID, prefix+"TargetID", "", ``)

	flags.StringVar(&v.Body, prefix+"Body", "", ``)

}

func printAddCommentRequest(v *pace.AddCommentRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistAddCommentRequest() {

	fmt.Print("-OrgID= ")

	fmt.Print("-TargetKind= ")

	fmt.Print("-TargetID= ")

	fmt.Print("-Body= ")

}

func addFlagsForPerson(flags *flag.FlagSet, prefix string, v *pace.Person) {

	flags.StringVar(&v.ID, prefix+"ID", "", ``)

	flags.StringVar(&v.Username, prefix+"Username", "", ``)

	flags.StringVar(&v.Name, prefix+"Name", "", ``)

	flags.StringVar(&v.PhotoURL, prefix+"PhotoURL", "", ``)

}

func printPerson(v *pace.Person) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistPerson() {

	fmt.Print("-ID= ")

	fmt.Print("-Username= ")

	fmt.Print("-Name= ")

	fmt.Print("-PhotoURL= ")

}

func addFlagsForFile(flags *flag.FlagSet, prefix string, v *pace.File) {

	flags.StringVar(&v.ID, prefix+"ID", "", `ID is the identifier for this file.`)

	flags.StringVar(&v.CTime, prefix+"CTime", "", `CTime is the time the file was uploaded.`)

	flags.StringVar(&v.Name, prefix+"Name", "", `Name is the name of the file.`)

	flags.StringVar(&v.Path, prefix+"Path", "", `Path is the path of the file.`)

	flags.StringVar(&v.ContentType, prefix+"ContentType", "", `ContentType is the type of the file.`)

	flags.StringVar(&v.FileType, prefix+"FileType", "", `FileType is the type of file.
Can be &#34;file&#34;, &#34;video&#34;, &#34;image&#34;, &#34;audio&#34; or &#34;screenshare&#34;.`)

	flags.StringVar(&v.DownloadURL, prefix+"DownloadURL", "", `DownloadURL URL which can be used to get the file.`)

	flags.StringVar(&v.ThumbnailURL, prefix+"ThumbnailURL", "", `ThumbnailURL is an optional thumbnail URL for this file.`)

	addFlagsForPerson(flags, "Author.", &v.Author)

}

func printFile(v *pace.File) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistFile() {

	fmt.Print("-ID= ")

	fmt.Print("-CTime= ")

	fmt.Print("-Name= ")

	fmt.Print("-Path= ")

	fmt.Print("-ContentType= ")

	fmt.Print("-FileType= ")

	fmt.Print("-Size= ")

	fmt.Print("-DownloadURL= ")

	fmt.Print("-ThumbnailURL= ")

	fmt.Print("-Author= ")

}

func addFlagsForComment(flags *flag.FlagSet, prefix string, v *pace.Comment) {

	flags.StringVar(&v.ID, prefix+"ID", "", ``)

	flags.StringVar(&v.CTime, prefix+"CTime", "", ``)

	flags.StringVar(&v.MTime, prefix+"MTime", "", ``)

	flags.StringVar(&v.Body, prefix+"Body", "", ``)

	flags.StringVar(&v.BodyHTML, prefix+"BodyHTML", "", ``)

	addFlagsForPerson(flags, "Author.", &v.Author)

	// []File not supported yet

}

func printComment(v *pace.Comment) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistComment() {

	fmt.Print("-ID= ")

	fmt.Print("-CTime= ")

	fmt.Print("-MTime= ")

	fmt.Print("-Body= ")

	fmt.Print("-BodyHTML= ")

	fmt.Print("-Author= ")

	fmt.Print("-Files= ")

}

func addFlagsForAddCommentResponse(flags *flag.FlagSet, prefix string, v *pace.AddCommentResponse) {

	addFlagsForComment(flags, "Comment.", &v.Comment)

	// skipping Error field (handled by Go client)

}

func printAddCommentResponse(v *pace.AddCommentResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistAddCommentResponse() {

	fmt.Print("-Comment= ")

	fmt.Print("-Error= ")

}

func addFlagsForRelatedCardsSummary(flags *flag.FlagSet, prefix string, v *pace.RelatedCardsSummary) {

}

func printRelatedCardsSummary(v *pace.RelatedCardsSummary) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistRelatedCardsSummary() {

	fmt.Print("-Total= ")

	fmt.Print("-Done= ")

	fmt.Print("-Progress= ")

}

func addFlagsForCard(flags *flag.FlagSet, prefix string, v *pace.Card) {

	flags.StringVar(&v.ID, prefix+"ID", "", `ID is the unique ID of the card within the org.`)

	flags.StringVar(&v.CTime, prefix+"CTime", "", ``)

	flags.StringVar(&v.MTime, prefix+"MTime", "", ``)

	flags.StringVar(&v.TeamID, prefix+"TeamID", "", ``)

	flags.StringVar(&v.Slug, prefix+"Slug", "", ``)

	flags.StringVar(&v.Title, prefix+"Title", "", ``)

	flags.StringVar(&v.Status, prefix+"Status", "", ``)

	addFlagsForPerson(flags, "Author.", &v.Author)

	flags.StringVar(&v.Body, prefix+"Body", "", ``)

	flags.StringVar(&v.BodyHTML, prefix+"BodyHTML", "", ``)

	// []string not supported yet

	// []Person not supported yet

	// []File not supported yet

	addFlagsForRelatedCardsSummary(flags, "RelatedCardsSummary.", &v.RelatedCardsSummary)

}

func printCard(v *pace.Card) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistCard() {

	fmt.Print("-ID= ")

	fmt.Print("-CTime= ")

	fmt.Print("-MTime= ")

	fmt.Print("-Order= ")

	fmt.Print("-TeamID= ")

	fmt.Print("-Slug= ")

	fmt.Print("-Title= ")

	fmt.Print("-Status= ")

	fmt.Print("-Author= ")

	fmt.Print("-Body= ")

	fmt.Print("-BodyHTML= ")

	fmt.Print("-Tags= ")

	fmt.Print("-TakenByCurrentUser= ")

	fmt.Print("-TakenByPeople= ")

	fmt.Print("-Files= ")

	fmt.Print("-RelatedCardsSummary= ")

}

func addFlagsForCreateCardRequest(flags *flag.FlagSet, prefix string, v *pace.CreateCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"OrgID", "", `OrgID is the org ID in which to create the card.`)

	flags.StringVar(&v.TeamID, prefix+"TeamID", "", `TeamID is the team ID in which to create the card.`)

	flags.StringVar(&v.Title, prefix+"Title", "", `Title is the title of the card.`)

	flags.StringVar(&v.ParentTargetKind, prefix+"ParentTargetKind", "", `ParentTargetKind is the kind of target to relate this card to (e.g. card or message)`)

	flags.StringVar(&v.ParentTargetID, prefix+"ParentTargetID", "", `ParentTargetID is the ID of the item to relate this new card to.`)

}

func printCreateCardRequest(v *pace.CreateCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistCreateCardRequest() {

	fmt.Print("-OrgID= ")

	fmt.Print("-TeamID= ")

	fmt.Print("-Title= ")

	fmt.Print("-ParentTargetKind= ")

	fmt.Print("-ParentTargetID= ")

}

func addFlagsForCreateCardResponse(flags *flag.FlagSet, prefix string, v *pace.CreateCardResponse) {

	addFlagsForCard(flags, "Card.", &v.Card)

	// skipping Error field (handled by Go client)

}

func printCreateCardResponse(v *pace.CreateCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistCreateCardResponse() {

	fmt.Print("-Card= ")

	fmt.Print("-Error= ")

}

func addFlagsForGetCardRequest(flags *flag.FlagSet, prefix string, v *pace.GetCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"OrgID", "", ``)

	flags.StringVar(&v.CardID, prefix+"CardID", "", ``)

}

func printGetCardRequest(v *pace.GetCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistGetCardRequest() {

	fmt.Print("-OrgID= ")

	fmt.Print("-CardID= ")

}

func addFlagsForGetCardResponse(flags *flag.FlagSet, prefix string, v *pace.GetCardResponse) {

	addFlagsForCard(flags, "Card.", &v.Card)

	// skipping Error field (handled by Go client)

}

func printGetCardResponse(v *pace.GetCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistGetCardResponse() {

	fmt.Print("-Card= ")

	fmt.Print("-Error= ")

}
