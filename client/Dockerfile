# Utiliser l'image de base Go
FROM golang:latest

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /go/src/app

# Copier les fichiers Go dans le conteneur
COPY . .

# Construire l'exécutable du serveur
RUN go build -o client client.go

# Exposer le port 50051 sur le conteneur
EXPOSE 50051

# Commande par défaut pour démarrer le serveur
CMD ["./client"]