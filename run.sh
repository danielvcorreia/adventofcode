year="2023"
day="01"
infile="input.puzzle"

while getopts "y:d:i:" op
do
    case $op in
        "y")
            year=$OPTARG
            ;;
        "d")
            day=$OPTARG
            ;;
        "i")
            infile=$OPTARG
            ;;
        default)
            echo "ERROR in parameters"
            ;;
    esac
done

shift $(($OPTIND-1))

cd ./$year/$day;
go run main.go < $infile;
